// Serve the seekr api using fiber
package api

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gofiber/template/html/v2"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	// "github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/seekr-osint/seekr/api/apierror"
	"github.com/seekr-osint/seekr/api/config"
	"github.com/seekr-osint/seekr/api/seekrauth"
	"github.com/swaggo/fiber-swagger"
)

// Serve the seekr api using fiber. called from the main package.
func Serve(config config.Config, fs embed.FS, db *gorm.DB, users seekrauth.Users) error {
	engine := html.NewFileSystem(http.FS(fs), ".html")

	app := fiber.New(fiber.Config{
		// ServerHeader: "Seekr",
		AppName: "Seekr",
		Views:   engine,
		// Global custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(apierror.GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})
	fav, err := fs.ReadFile("web/images/favicon.ico")
	if err != nil {
		return err
	}
	app.Use(favicon.New(favicon.Config{
		Data: fav,
	}))

	app.Use(basicauth.New(basicauth.Config{
		Users: users.ToMap(),
	}))
	// Logging remote IP and Port
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// app.Use(helmet.New(helmet.Config{
	// 	XSSProtection:             "0",
	// 	ContentTypeNosniff:        "sniff",
	// 	XFrameOptions:             "SAMEORIGIN",
	// 	ReferrerPolicy:            "no-referrer",
	// 	CrossOriginEmbedderPolicy: "require-corp",
	// 	CrossOriginOpenerPolicy:   "same-origin",
	// 	CrossOriginResourcePolicy: "same-origin",
	// 	OriginAgentCluster:        "?1",
	// 	XDNSPrefetchControl:       "off",
	// 	XDownloadOptions:          "noopen",
	// 	XPermittedCrossDomain:     "none",
	// }))

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	web := app.Group("/web")
	web.Use("/", filesystem.New(filesystem.Config{
		Root:   http.FS(fs),
		Browse: true,
		Index:  "index.html",
		// NotFoundFile: "404.html",
		MaxAge:     1,
		PathPrefix: "web",
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1", func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v1")
		return c.Next()
	})

	v1.Get("/people/:id", FiberHandler(GetPerson, db))
	v1.Get("/people", FiberHandler(GetPeople, db))
	v1.Delete("/people/:id", FiberHandler(DeletePerson, db))
	v1.Patch("/people/:id", FiberHandler(PatchPerson, db))
	v1.Post("/people", FiberHandler(PostPerson, db))

	v1.Get("/restart", FiberHandler(Restart, db))
	v1.Post("/detect-language", FiberHandler(DetectLanguage, db))
	v1.Get("/scanAccounts/:username", FiberHandler(Restart, db))

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Seekr Metrics Page"}))
	for _, route := range app.GetRoutes(true) {
		fmt.Printf("%s\t-> %s\n", route.Method, route.Path)
	}

	return app.Listen(config.Address())
}

// Handler for api endpoints.
// Limiting the DataBase to only include entries of the user making the request.
func FiberHandler(fn func(*fiber.Ctx, *gorm.DB) error, db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		localdb := UserDB(c.Locals("username").(string), db)
		return fn(c, localdb)
	}
}

// returning the DataBase where entries a given username is the owner of.
//
// It is used by FiberHandler() for multi user support.
func UserDB(username string, db *gorm.DB) *gorm.DB {
	return db.Where("owner = ?", username)
}

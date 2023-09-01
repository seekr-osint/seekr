package api

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/seekr-osint/seekr/api/language"
	"github.com/seekr-osint/seekr/api/person"
	"github.com/seekr-osint/seekr/api/restart"
)

func DetectLanguage(c *fiber.Ctx, db *gorm.DB) error {
	var text struct {
		Text string `json:"text"`
	}

	body := c.Body()
	err := json.Unmarshal(body, &text)
	if err != nil {
		return c.Status(503).SendString(err.Error())
	}
	lang := language.DetectLanguage(text.Text)
	return c.Status(200).JSON(lang)
}
func Restart(c *fiber.Ctx, db *gorm.DB) error {
	fmt.Printf("Restarting...\n")
	err := restart.RestartBinary()
	if err != nil {
		return err
	}
	return nil
}
func GetPerson(c *fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id")
	var person person.Person

	result := db.Find(&person, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&person)
}

func GetPeople(c *fiber.Ctx, db *gorm.DB) error {
	var people []person.Person
	db.Find(&people)
	res := map[uint]person.Person{}
	for _, person := range people {
		res[person.ID] = person
	}
	return c.Status(200).JSON(res)
}

func DeletePerson(c *fiber.Ctx, db *gorm.DB) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	result := db.Delete(&person.Person{}, uint(id))

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
func PatchPerson(c *fiber.Ctx, db *gorm.DB) error {
	person := person.Person{}
	id := c.Params("id")

	body := c.Body()
	err := json.Unmarshal(body, &person)
	if err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&person)
	return c.Status(200).JSON(person)
}

func PostPerson(c *fiber.Ctx, db *gorm.DB) error {
	person := person.Person{}

	body := c.Body()
	err := json.Unmarshal(body, &person)
	if err != nil {
		return c.Status(503).SendString(err.Error())
	}
	person.Owner = c.Locals("username").(string)
	if err = person.Validate(nil); err != nil {
		return err
	}

	db.Create(&person)
	return c.Status(201).JSON(person)
}


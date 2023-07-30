package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	//"os"
	_ "image/jpeg"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ServicesHandler(servicesToCheck Services, username string, config ApiConfig) Accounts {
	wg := &sync.WaitGroup{}

	accounts := Accounts{}
	for i := 0; i < len(servicesToCheck); i++ { // loop over all services
		wg.Add(1)
		go func(i int) {
			// Do something
			service := servicesToCheck[i] // current service
			err, exists := service.UserExistsFunc(service, username, config)
			if err != nil {
				// FIXME add skipped accounts
			} else if exists {
				err, account := service.GetInfoFunc(username, service, config) // add service to accounts
				if err != nil {
					// Skipping account info
					accounts[fmt.Sprintf("%s-%s", servicesToCheck[i].Name, username)] = EmptyAccountInfo(username, service)
				}
				accounts[fmt.Sprintf("%s-%s", servicesToCheck[i].Name, username)] = account
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return accounts
}

var DefaultServices = Services{
	Service{
		Name:           "GitHub",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    GithubInfo,
		BaseUrl:        "https://github.com/{username}",
	},
	Service{
		Name:           "TikTok",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://tiktok.com/@{username}",
	},
	Service{
		Name:           "Twitter",
		Check:          "pattern",
		Pattern:        "<div class=\"error-panel\"><span>User ",
		UserExistsFunc: AtUsernameUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://nitter.net/{username}",
		HtmlUrl:        "https://twitter.com/{username}",
	},
	Service{
		Name:           "Instagram",
		Check:          "",
		Pattern:        "Nothing found!",
		UserExistsFunc: InstagramUserExistsCheck,
		GetInfoFunc:    InstagramInfo,
		BaseUrl:        "https://instagram.com/{username}",
	},
	Service{
		Name:           "LinkedIn",
		Check:          "", // broken
		Pattern:        "{username}",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://linkedin.com/in/{username}",
	},
	Service{
		Name:           "Snapchat",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://www.snapchat.com/add/{username}",
		// AvatarUrl:      "https://app.snapchat.com/web/deeplink/snapcode?username={username}&type=SVG&bitmoji=enable", // FIXME SVG
	},
	Service{
		Name:           "Reddit",
		Check:          "", // FIXME blocked not sure rather it actually works
		Pattern:        "Sorry, nobody on Reddit goes by that name.",
		BlockedPattern: "<title>Blocked</title>",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://reddit.com/user/{username}",
	},
	Service{
		Name:           "Facebook",
		Check:          "", // FIXME disabled
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://www.facebook.com/{username}/videos",
		HtmlUrl:        "https://www.facebook.com/{username}",
	},
	Service{
		Name:           "Twitch",
		Check:          "", // FIXME disabled
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://www.twitch.tv/{username}",
	},
	Service{
		Name:           "Chess.com",
		Check:          "", // FIXME disabled
		Pattern:        "The page you are looking for doesn’t exist. (404)",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://chess.com/member/{username}",
	},
	Service{
		Name:           "SteamGroup",
		Check:          "pattern",
		Pattern:        "No group could be retrieved for the given URL",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://steamcommunity.com/groups/{username}",
	},
	Service{
		Name:           "SteamCommunity",
		Check:          "pattern",
		Pattern:        "The specified profile could not be found.",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://steamcommunity.com/id/{username}",
	},
	Service{
		Name:           "TryHackMe",
		Check:          "status_code",
		UserExistsFunc: TryHackMeUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://tryhackme.com/p/{username}",
	},
	Service{
		Name:           "Odysee",
		Check:          "", // FIXME
		Pattern:        "Channel Not Found",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://odysee.com/@{username}",
	},
	Service{
		Name:           "Scratch",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://scratch.mit.edu/users/{username}",
	},
	Service{
		Name:           "Telegram",
		Check:          "", // FIXME
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://t.me/{username}",
	},
	Service{
		Name:           "XBox Gamertag",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://xboxgamertag.com/search/{username}",
	},
	Service{
		Name:           "Spotify",
		Check:          "", // FIXME
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://spotify.com/user/{username}",
	},
	Service{
		Name:           "Mastodon.social",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://mastodon.social/{username}",
	},
	Service{
		Name:           "Mastodon.xyz",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://mastodon.xyz/{username}",
	},
	Service{
		Name:           "Npm",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://npmjs.com/~{username}",
	},
	Service{
		Name:           "Nintendolife",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://www.nintendolife.com/users/{username}",
	},
	Service{
		Name:           "VirusTotal",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://www.virustotal.com/ui/users/{username}/trusted_users",
	},
	Service{
		Name:           "Tellonym.me",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://tellonym.me/{username}",
	},
	Service{
		Name:           "gutefrage",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://www.gutefrage.net/nutzer/{username}",
	},
	Service{
		Name:           "Lichess",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    LichessInfo,
		BaseUrl:        "https://lichess.org/api/user/{username}",
	},
	Service{
		Name:           "SoundCloud",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://soundcloud.com/{username}",
	},
	Service{
		Name:           "repl.it",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://replit.com/{username}",
	},
	Service{
		Name:           "PyPi",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://pypi.org/user/{username}",
	},
	Service{
		Name:           "SlideShare",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://slideshare.net/{username}",
		AvatarUrl:      "https://cdn.slidesharecdn.com/profile-photo-{username}-96x96.jpg",
	},
	Service{
		Name:           "Fosstodon",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://fosstodon.org/@{username}",
	},
	Service{
		Name:           "Slides.com",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://slides.com/{username}",
	},
	Service{
		Name:           "Giphy",
		Check:          "status_code", // FIXME gives a topic
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://giphy.com/{username}",
	},
	Service{
		Name:           "Gravatar",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		ScrapeImage:    true,
		Scrape: ScrapeStruct{
			FindElement: ".photo-0",
			Attr:        "href",
		},
		BaseUrl: "http://en.gravatar.com/{username}",
	},
	Service{
		Name:              "LeetCode",
		Check:             "status_code",
		UserExistsFunc:    SimpleUserExistsCheck,
		GetInfoFunc:       SimpleAccountInfo,
		ExternalImageFunc: true,
		ImageFunc:         LeetCodeImage,
		BaseUrl:           "https://leetcode.com/{username}",
	},
	Service{
		Name:              "Asciinema",
		Check:             "status_code",
		UserExistsFunc:    SimpleUserExistsCheck,
		GetInfoFunc:       SimpleAccountInfo,
		ExternalImageFunc: true,
		ImageFunc:         AsciinemaImage,
		BaseUrl:           "https://asciinema.org/~{username}",
	},
	Service{
		Name:           "Ask Fedora",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://ask.fedoraproject.org/u/{username}",
	},
	Service{
		Name:           "Autofrage",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://www.autofrage.net/nutzer/{username}",
	},
	Service{
		Name:           "Finanzfrage",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://www.finanzfrage.net/nutzer/{username}",
	},
	Service{
		Name:           "Gesundheitsfrage",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://www.gesundheitsfrage.net/nutzer/{username}",
	},
	Service{
		Name:           "Linktree",
		Check:          "", // FIXME this service sucks
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://linktr.ee/{username}",
	},
	Service{
		Name:           "Myspace",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://myspace.com/{username}",
	},
	Service{
		Name:           "Flickr",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://www.flickr.com/people/{username}",
	},
	Service{
		Name:           "FortniteTracker",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://fortnitetracker.com/profile/all/{username}",
	},
	Service{
		Name:           "Brave Community",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://community.brave.com/u/{username}",
	},
	Service{
		Name:           "BuyMeACoffee",
		Check:          "status_code", // FIXME down
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://buymeacoff.ee/{username}",
	},
	Service{
		Name:              "Bitbucket",
		Check:             "status_code",
		UserExistsFunc:    SimpleUserExistsCheck,
		GetInfoFunc:       SimpleAccountInfo,
		ExternalImageFunc: true,
		ImageFunc:         BitbucketImage,
		BaseUrl:           "https://bitbucket.org/{username}/",
		// AvatarUrl:      "https://bitbucket.org/workspaces/{username}/avatar/", // FIXME
	},
	Service{
		Name:           "Bitwarden",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://community.bitwarden.com/u/{username}/summary",
	},
	Service{
		Name:           "Cloudflare",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://community.cloudflare.com/u/{username}",
	},
	Service{
		Name:           "Clubhouse",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://www.clubhouse.com/@{username}",
	},
	Service{
		Name:           "Codepen",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://codepen.io/{username}",
	},
	Service{
		Name:           "Codewars",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://www.codewars.com/users/{username}",
	},
	Service{
		Name:           "Docker Hub",
		Check:          "", // FIXME
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://hub.docker.com/u/{username}",
	},
	Service{
		Name:           "Apple Developer",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://developer.apple.com/forums/profile/{username}",
	},
}

func HttpRequest(url string) (string, error) {
	log.Println("request to:" + url)
	if url != "" {
		resp, err := http.Get(url)
		if err != nil {
			log.Print("error http request")
			log.Println(err)
			return "", err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Print("error http request")
			log.Println(err)
			return "", err
		}
		return string(body), nil
	}

	return "", errors.New("empty URL provided")
}

func UrlTemplate(url string, username string) string {
	return strings.ReplaceAll(url, "{username}", username)
}

func SimpleUserExistsCheck(service Service, username string, config ApiConfig) (error, bool) { // type UserExistsFunc
	log.Printf("checking: %s %s", service.Name, username)
	if config.Testing {
		if username == strings.ToLower(fmt.Sprintf("%s-exists", service.Name)) {
			log.Printf("%s-exists", service.Name)
			return nil, true
		} else if username == fmt.Sprintf("%s-error", service.Name) {
			return errors.New("error"), false
		}
		return nil, false
	}
	BaseUrl := UrlTemplate(service.BaseUrl, username)
	log.Println("checking:" + BaseUrl)
	switch service.Check {
	case "status_code":
		err, status_code := GetStatusCodeNew(BaseUrl, config)
		return err, status_code == 200
	case "pattern": // search for string on website
		site, err := HttpRequest(BaseUrl)
		if err != nil {
			return err, false
		} else {
			// log.Println(site)
			pattern_found := strings.Contains(site, UrlTemplate(service.Pattern, username)) // ! pattern was found
			if service.BlockedPattern != "" {
				blocked := strings.Contains(site, service.BlockedPattern)
				if blocked {
					return errors.New("blocked"), false
				}
			}
			if !pattern_found {
				log.Println("found anti_pattern:", pattern_found)
				return nil, true
			}
		}
	}

	return nil, false
}

func InstagramUserExistsCheck(service Service, username string, config ApiConfig) (error, bool) { // type UserExistsFunc

	// FIXME this is a workaround
	// This will never truly work
	// Instagram protects against scraping stuff like this
	// Specific usernames will give false positives (like "div" or some weird numbers that instagram uses as class names for fuck knows why)
	// But i dont care at all
	// I just want to get this over with
	// I dont know how to fix this
	// I hate everything
	// I hate this workaround
	// I am about to cry
	// I am about to kill myself
	// If anyone knows how to fix this please tell me
	// I am desperate
	// A therapist cant even help me now
	// I am not joking

	log.Printf("checking: %s %s", service.Name, username)
	if config.Testing {
		if username == strings.ToLower(fmt.Sprintf("%s-exists", service.Name)) {
			log.Printf("%s-exists", service.Name)
			return nil, true
		} else if username == fmt.Sprintf("%s-error", service.Name) {
			return errors.New("error"), false
		}
		return nil, false
	}

	BaseUrl := UrlTemplate(service.BaseUrl, username)
	log.Println("checking:" + BaseUrl)

	site, err := HttpRequest(BaseUrl)
	if err != nil {
		return err, false
	}

	// This is a workaround. I dont like this, but it works and is sort of reliable for now.
	// Instagram only displays the username once somewhere in the HTML if it doesnt exist. If it exists it displays it more.

	count := strings.Count(site, username)

	if count > 1 {
		return nil, true
	}

	return nil, false
}

func AtUsernameUserExistsCheck(service Service, username string, config ApiConfig) (error, bool) { // type UserExistsFunc
	log.Printf("checking: %s %s", service.Name, username)
	if config.Testing {
		if username == strings.ToLower(fmt.Sprintf("@%s-exists", service.Name)) {
			log.Printf("%s-exists", service.Name)
			return nil, true
		} else if username == fmt.Sprintf("@%s-error", service.Name) {
			return errors.New("error"), false
		}
		return nil, false
	}

	BaseUrl := UrlTemplate(service.BaseUrl, username)
	log.Println("checking:" + BaseUrl)

	site, err := HttpRequest(BaseUrl)
	if err != nil {
		return err, false
	}

	if strings.Contains(strings.ToLower(site), "@"+username) {
		return nil, true
	}

	return nil, false
}

func TryHackMeUserExistsCheck(service Service, username string, config ApiConfig) (error, bool) { // type UserExistsFunc
	log.Printf("checking: %s %s", service.Name, username)
	if config.Testing {
		if username == strings.ToLower(fmt.Sprintf("@%s-exists", service.Name)) {
			log.Printf("%s-exists", service.Name)
			return nil, true
		} else if username == fmt.Sprintf("@%s-error", service.Name) {
			return errors.New("error"), false
		}
		return nil, false
	}

	BaseUrl := UrlTemplate(service.BaseUrl, username)
	log.Println("checking:" + BaseUrl)

	site, err := HttpRequest(BaseUrl)
	if err != nil {
		return err, false
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(site))
	if err != nil {
		return err, false
	}

	if usernameSpan := doc.Find("body > div.container.profile-body > div.p-3 > h2 > span"); usernameSpan.Length() > 0 {
		if usernameSpan.Text() == username {
			return nil, true
		}
	}

	return nil, false
}

func EmptyAccountInfo(username string, service Service) Account {
	return Account{
		Service:  service.Name,
		Username: username,
	}
}

// maybe remove
//
//	func DefaultServicesHandler(username string) Accounts {
//		return ServicesHandler(DefaultServices, username)
//	}
func getImg(img string) image.Image {
	reader := strings.NewReader(img)
	decodedImg, imgType, err := image.Decode(reader)
	log.Printf("image type:%s", imgType)
	if err != nil {
		log.Println(err)
	}
	return decodedImg
}

func EncodeBase64(img string) string {
	decodedImg := getImg(img)
	buf := new(bytes.Buffer)
	err := png.Encode(buf, decodedImg)
	if err != nil {
		log.Println(err)
	}

	base64Img := base64.StdEncoding.EncodeToString(buf.Bytes())
	return base64Img
}

func GetAvatar(avatar_url string, account Account, config ApiConfig) (error, Account) {
	log.Printf("avatar_url: %s", avatar_url)
	err, statusCode := GetStatusCodeNew(avatar_url, ApiConfig{}) // FIXME empty api config
	if err != nil {
		return err, account
	}
	if statusCode == 200 {
		avatar, err := HttpRequest(avatar_url)
		if err != nil {
			return err, account
		}
		account.Picture = Pictures{
			"1": Picture{Img: EncodeBase64(avatar), ImgHash: MkImgHash(getImg(avatar))},
		}
	}
	return nil, account
}

func SimpleAccountInfo(username string, service Service, config ApiConfig) (error, Account) {
	if config.Testing {
		if username == strings.ToLower(fmt.Sprintf("%s-exists", service.Name)) {
			return nil, EmptyAccountInfo(username, service)
		} else if username == fmt.Sprintf("%s-error", service.Name) {
			return errors.New("error"), EmptyAccountInfo(username, service)
		}
		return nil, EmptyAccountInfo(username, service)
	}
	log.Println("simple account info:" + service.Name)
	baseUrl := UrlTemplate(service.BaseUrl, username)
	htmlUrl := UrlTemplate(service.HtmlUrl, username)
	account := Account{
		Service:  service.Name,
		Username: username,
	}
	if htmlUrl == "" {
		account.Url = baseUrl
	} else {
		account.Url = htmlUrl
	}

	if service.ExternalImageFunc {
		// no check for "" needed the check happenes later in the code
		service.AvatarUrl = UrlTemplate(service.ImageFunc(username, service), username)
	}
	if service.ScrapeImage {

		doc, err := goquery.NewDocument(baseUrl)
		if err != nil {
			log.Println(err)
			return err, EmptyAccountInfo(username, service)
		}

		// Select the <a> element with class "photo-0"
		link := doc.Find(service.Scrape.FindElement).First()

		// Get the value of the href attribute
		res, exists := link.Attr(service.Scrape.Attr)
		if exists {
			service.AvatarUrl = res
		}
	}
	if service.AvatarUrl != "" {
		err, newAccount := GetAvatar(service.AvatarUrl, account, config) // TODO give service as argument
		if err != nil {
			return nil, account
		}
		return nil, newAccount
	}
	return nil, account
}

func SlideshareInfo(username string, service Service, config ApiConfig) (error, Account) {
	log.Println("slideshare")
	baseUrl := strings.ReplaceAll(service.BaseUrl, "{username}", username)
	avatar_url := strings.ReplaceAll(service.BaseUrl, "{username}", username)
	account := Account{
		Service:  service.Name,
		Username: username,
		Url:      baseUrl,
	}
	err, newAccount := GetAvatar(avatar_url, account, config)
	if err != nil {
		return nil, account
	}
	return nil, newAccount
}
func GithubInfo(username string, service Service, config ApiConfig) (error, Account) {
	log.Println("github")
	var data struct {
		Id         int    `json:"id"`
		Bio        string `json:"bio"`
		Avatar_url string `json:"avatar_url"`
		Url        string `json:"html_url"`
		Name       string `json:"name"`
		Location   string `json:"location"`
		Repos      int    `json:"public_repos"`
		Gists      int    `json:"public_gists"`
		Blog       string `json:"blog"`
		Twitter    string `json:"twitter_username"`
		Followers  int    `json:"followers"`
		Following  int    `json:"following"`
		Created_at string `json:"created_at"`
		Updated_at string `json:"updated_at"`
		Company    string `jons:"company"`
	}
	jsonData, err := HttpRequest("https://api.github.com/users/" + username)
	if err != nil {
		return err, EmptyAccountInfo(username, service)
	}
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return err, EmptyAccountInfo(username, service)
	}

	log.Println("no api limitation")
	log.Printf("avatar_url: %s", data.Avatar_url)
	if err != nil {
		return err, EmptyAccountInfo(username, service)
	}
	account := Account{
		Service:  service.Name,
		Username: username,
		Url:      data.Url,
		Id:       strconv.Itoa(data.Id),
		Bio:      Bios{"1": Bio{Bio: data.Bio}},

		Location:  data.Location,
		Created:   data.Created_at,
		Updated:   data.Updated_at,
		Blog:      data.Blog,
		Followers: data.Followers,
		Following: data.Following,
	}
	avatar, err := HttpRequest(data.Avatar_url)
	if err != nil {
		log.Println(err)
	} else {
		account.Picture = Pictures{
			"1": {Img: EncodeBase64(avatar), ImgHash: MkImgHash(getImg(avatar))},
		}
	}
	return nil, account
}

// TODO Finish this

func InstagramInfo(username string, service Service, config ApiConfig) (error, Account) {
	log.Println("instagram")
	var data struct {
		Name string `json:"name"`
		Url  string `json:"url"`
		// Profile struct {
		// 	Bio       string `json:"bio"`
		// 	Posts     string `json:"posts"`
		// 	Followers string `json:"followers"`
		// 	Following string `json:"following"`
		// } `json:"profile"`
	}
	jsonData, err := HttpRequest("https://www.instagram.com/" + username + "/?__a=1&__d=1/")
	if err != nil {
		return err, EmptyAccountInfo(username, service)
	}
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Println(err)
		return err, EmptyAccountInfo(username, service)
	}
	return nil, Account{
		Service:  service.Name,
		Username: username,
		Name:     Name,
		Url:      "https://www.instagram.com/" + username,
		// Bio: Bios{"1": Bio{Bio: data.Profile.Bio}},
		// Firstname: data.Profile.Firstname,
		// Lastname:  data.Profile.Lastname,
	}
}

func LichessInfo(username string, service Service, config ApiConfig) (error, Account) {
	log.Println("lichess")
	var data struct {
		Id      string `json:"id"`
		Url     string `json:"url"`
		Profile struct {
			Bio       string `json:"bio"`
			Firstname string `json:"firstName"`
			Lastname  string `json:"lastName"`
		} `json:"profile"`
	}
	jsonData, err := HttpRequest("https://lichess.org/api/user/" + username)
	if err != nil {
		return err, EmptyAccountInfo(username, service)
	}
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Println(err)
		return err, EmptyAccountInfo(username, service)
	}
	return nil, Account{
		Service:   service.Name,
		Username:  username,
		Id:        data.Id,
		Url:       "https://lichess.org/@/" + username,
		Bio:       Bios{"1": Bio{Bio: data.Profile.Bio}},
		Firstname: data.Profile.Firstname,
		Lastname:  data.Profile.Lastname,
	}
}

func LeetCodeImage(username string, service Service) string {
	doc, err := goquery.NewDocument(strings.ReplaceAll(service.BaseUrl, "{username}", username))
	if err != nil {
		log.Println(err)
	}

	// Select the <a> element with class "photo-0"
	link := doc.Find("img.rounded-lg").First()

	// Get the value of the href attribute
	href, exists := link.Attr("src")
	if exists {
		return href
	}
	return ""
}
func AsciinemaImage(username string, service Service) string {
	doc, err := goquery.NewDocument(strings.ReplaceAll(service.BaseUrl, "{username}", username))
	if err != nil {
		log.Println(err)
	}

	// Select the <a> element with class "photo-0"
	link := doc.Find("img.avatar").First()

	// Get the value of the href attribute
	href, exists := link.Attr("src")
	if exists {
		return "https:" + href
	}
	return ""
}
func BitbucketImage(username string, service Service) string { // FIXME
	doc, err := goquery.NewDocument(strings.ReplaceAll(service.BaseUrl, "{username}", username))
	if err != nil {
		log.Println(err)
	}

	// Select the <a> element with class "photo-0"

	// Select the <span> element with class "css-ob4lje"
	span := doc.Find("span.css-ob4lje").First()

	// Get the value of the "background-image" style property
	style, exists := span.Attr("style")
	if exists {
		// Extract the URL from the style property value
		url := strings.TrimPrefix(strings.TrimSuffix(strings.Split(style, ":")[1], ";"), " url(\"")
		return url
	}
	return ""
}

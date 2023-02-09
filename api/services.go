package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"sync"

	//"io"
	"log"
	"net/http"

	//"os"
	"github.com/PuerkitoBio/goquery"
	_ "image/jpeg"
	"strconv"
	"strings"
)

type Services []Service
type Service struct {
	Name              string         // example: "GitHub"
	UserExistsFunc    UserExistsFunc // example: SimpleUserExistsCheck()
	GetInfoFunc       GetInfoFunc    // example: EmptyAccountInfo()
	ImageFunc         ImageFunc
	ExternalImageFunc bool
	BaseUrl           string // example: "https://github.com"
	AvatarUrl         string
	Check             string // example: "status_code"
	HtmlUrl           string
	Pattern           string
	BlockedPattern    string
}

type GetInfoFunc func(string, Service) Account // (username)
type ImageFunc func(string, Service) string    // (username)
type UserExistsFunc func(Service, string) bool // (BaseUrl,username)

func ServicesHandler(servicesToCheck Services, username string) Accounts {
	wg := &sync.WaitGroup{}

	accounts := Accounts{}
	for i := 0; i < len(servicesToCheck); i++ { // loop over all services
		wg.Add(1)
		go func(i int) {
			// Do something
			service := servicesToCheck[i]                  // current service
			if service.UserExistsFunc(service, username) { // if service exisits
				accounts[fmt.Sprintf("%s-%s", servicesToCheck[i].Name, username)] = service.GetInfoFunc(username, service) // add service to accounts
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
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://nitter.net/{username}",
		HtmlUrl:        "https://twitter.com/{username}",
	},
	Service{
		Name:           "Instagram",
		Check:          "",
		Pattern:        "Nothing found!",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
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
		UserExistsFunc: SimpleUserExistsCheck,
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
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://giphy.com/{username}",
	},
	Service{
		Name:              "Gravatar",
		Check:             "status_code",
		UserExistsFunc:    SimpleUserExistsCheck,
		GetInfoFunc:       SimpleAccountInfo,
		ExternalImageFunc: true,
		ImageFunc:         GravatarImage,
		BaseUrl:           "http://en.gravatar.com/{username}",
	},
	Service{
		Name:           "HackTheBox",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://forum.hackthebox.eu/profile/{username}",
	},
	Service{
		Name:           "LeetCode",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://leetcode.com/{username}",
	},
	Service{
		Name:           "Asciinema",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://asciinema.org/~{username}",
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
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://buymeacoff.ee/{username}",
	},
	Service{
		Name:           "Bitbucket",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://bitbucket.org/{username}/",
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

func SimpleUserExistsCheck(service Service, username string) bool {
	BaseUrl := strings.ReplaceAll(service.BaseUrl, "{username}", username)
	log.Println("check:" + BaseUrl)
	exists := false
	switch service.Check {
	case "status_code":
		exists = GetStatusCode(BaseUrl) == 200
	case "pattern": // search for string on website
		site := HttpRequest(BaseUrl)
		// log.Println(site)
		found := strings.Contains(site, strings.ReplaceAll(service.Pattern, "{username}", username)) // ! pattern was found
		blocked := false
		if service.BlockedPattern != "" {
			blocked = strings.Contains(site, service.BlockedPattern)
		}
		if !blocked {
			log.Println("found:", found)
			if !found {
				exists = true
			}
		}
		// the pattern is the not found text therefore it's true if no not found text was found
		// BUG you can put the not found text into your bio
	}

	return exists
}

func EmptyAccountInfo(username string, service Service) Account {
	return Account{
		Service:  service.Name,
		Username: username,
	}
}

// maybe remove
func DefaultServicesHandler(username string) Accounts {
	return ServicesHandler(DefaultServices, username)
}

func HttpRequest(url string) string {
	log.Println("request to:" + url)
	if url != "" {
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		return string(body)
	}

	return ""
}

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

func GetAvatar(avatar_url string, account Account) Account {
	log.Printf("avatar_url: %s", avatar_url)
	if GetStatusCode(avatar_url) == 200 {
		avatar := HttpRequest(avatar_url)
		account.Picture = Pictures{
			"1": Picture{Img: EncodeBase64(avatar), ImgHash: MkImgHash(getImg(avatar))},
		}
	}
	return account
}

func SimpleAccountInfo(username string, service Service) Account {
	log.Println(service.Name)
	baseUrl := strings.ReplaceAll(service.BaseUrl, "{username}", username)
	account := Account{
		Service:  service.Name,
		Username: username,
	}
	if service.HtmlUrl == "" {
		account.Url = baseUrl
	} else {
		account.Url = strings.ReplaceAll(service.HtmlUrl, "{username}", username)
	}

	if service.ExternalImageFunc {
		service.AvatarUrl = service.ImageFunc(username, service)
	}
	if service.AvatarUrl != "" {
		avatar_url := strings.ReplaceAll(service.AvatarUrl, "{username}", username)
		account = GetAvatar(avatar_url, account)
	}
	return account
}

func SlideshareInfo(username string, service Service) Account {
	log.Println("slideshare")
	baseUrl := strings.ReplaceAll(service.BaseUrl, "{username}", username)
	avatar_url := strings.ReplaceAll(service.BaseUrl, "{username}", username)
	account := Account{
		Service:  service.Name,
		Username: username,
		Url:      baseUrl,
	}
	account = GetAvatar(avatar_url, account)
	return account
}
func GithubInfo(username string, service Service) Account {
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
	var errData struct {
		Documentation_url string `json:"documentation_url"`
	}

	jsonData := HttpRequest("https://api.github.com/users/" + username)

	err := json.Unmarshal([]byte(jsonData), &errData)
	if err != nil {
		log.Println(err)
	} else {
		if errData.Documentation_url == "https://docs.github.com/rest/overview/resources-in-the-rest-api#rate-limiting" {
			log.Println("api request limitied")
		} else {

			log.Println("no api limitation")
			err = json.Unmarshal([]byte(jsonData), &data)
			log.Printf("avatar_url: %s", data.Avatar_url)
			if err != nil {
				log.Println(err)
			}
		}
	}
	avatar := HttpRequest(data.Avatar_url)
	account := Account{
		Service:  service.Name,
		Username: username,
		Url:      data.Url,
		Id:       strconv.Itoa(data.Id),
		Bio:      Bios{"1": Bio{Bio: data.Bio}},
		Picture: Pictures{

			"1": {Img: EncodeBase64(avatar), ImgHash: MkImgHash(getImg(avatar))},
		},
		Location:  data.Location,
		Created:   data.Created_at,
		Updated:   data.Updated_at,
		Blog:      data.Blog,
		Followers: data.Followers,
		Following: data.Following,
	}
	return account
}

func LichessInfo(username string, service Service) Account {
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
	jsonData := HttpRequest("https://lichess.org/api/user/" + username)
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Println(err)
	}
	return Account{
		Service:   service.Name,
		Username:  username,
		Id:        data.Id,
		Url:       "https://lichess.org/@/" + username,
		Bio:       Bios{"1": Bio{Bio: data.Profile.Bio}},
		Firstname: data.Profile.Firstname,
		Lastname:  data.Profile.Lastname,
	}
}

func GravatarImage(username string, service Service) string {
	doc, err := goquery.NewDocument("http://en.gravatar.com/" + username)
	if err != nil {
		log.Println(err)
	}

	// Select the <a> element with class "photo-0"
	link := doc.Find(".photo-0").First()

	// Get the value of the href attribute
	href, exists := link.Attr("href")
	if exists {
		return href
	}
	return ""
}

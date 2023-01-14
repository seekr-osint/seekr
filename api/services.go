package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image"
	"image/png"
	"io/ioutil"
	//"io"
	"log"
	"net/http"
	//"os"
	_ "image/jpeg"
	"strconv"
	"strings"
)

var DefaultServices = Services{
	Service{
		Name:           "GitHub",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    GithubInfo,
		BaseUrl:        "https://github.com/{username}",
	},
	Service{
		Name:           "Facebook",
		Check:          "", // FIXME
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://www.facebook.com/{username}/videos",
		HtmlUrl:        "https://www.facebook.com/{username}",
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
		Name:           "SlideShare",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://slideshare.net/{username}",
		AvatarUrl:      "https://cdn.slidesharecdn.com/profile-photo-{username}-96x96.jpg",
	},
	Service{
		Name:           "Slides.com",
		Check:          "status_code",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    SimpleAccountInfo,
		BaseUrl:        "https://slides.com/{username}",
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
	Service{ // broken
		Name:           "Reddit",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    RedditInfo,
		BaseUrl:        "https://api.reddit.com/user/{username}",
	},
}

type Services []Service
type Service struct {
	Name           string         // example: "GitHub"
	UserExistsFunc UserExistsFunc // example: SimpleUserExistsCheck()
	GetInfoFunc    GetInfoFunc    // example: EmptyAccountInfo()
	BaseUrl        string         // example: "https://github.com"
	AvatarUrl      string
	Check          string // example: "status_code"
	HtmlUrl        string
}

// type Accounts map[string]Account
type Accounts []Account
type Account struct {
	Service   string   `json:"service"`  // example: GitHub
	Id        string   `json:"id"`       // example: 1224234
	Username  string   `json:"username"` // example: 9glenda
	Url       string   `json:"url"`      // example: https://github.com/9glenda
	Picture   []string `json:"profilePicture"`
	ImgHash   []uint64 `json:"imgHash"`
	Bio       []string `json:"bio"`       // example: pro hacka
	Firstname string   `json:"firstname"` // example: Glenda
	Lastname  string   `json:"lastname"`  // example: Belov
	Location  string   `json:"location"`  // example: Moscow
	Created   string   `json:"created"`   // example: 2020-07-31T13:04:48Z
	Updated   string   `json:"updated"`
	Blog      string   `json:"blog"`
	Followers int      `json:"followers"`
	Following int      `json:"following"`
}

type GetInfoFunc func(string, Service) Account // (username)
type UserExistsFunc func(Service, string) bool // (BaseUrl,username)

func SimpleUserExistsCheck(service Service, username string) bool {
	BaseUrl := strings.ReplaceAll(service.BaseUrl, "{username}", username)
	log.Println("check:" + BaseUrl)
	//log.Println(GetStatusCode(BaseUrl))
	exists := false
	if service.Check == "status_code" {
		exists = GetStatusCode(BaseUrl) == 200
	}
	return exists
}

func EmptyAccountInfo(username string, service Service) Account {
	return Account{
		Service:  service.Name,
		Username: username,
		Bio:      nil,
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

func ServicesHandler(servicesToCheck Services, username string) Accounts {
	var accounts Accounts
	for i := 0; i < len(servicesToCheck); i++ {
		service := servicesToCheck[i]
		if service.UserExistsFunc(service, username) {
			accounts = append(accounts, service.GetInfoFunc(username, service))
		}
	}
	return accounts
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
		account.Picture = []string{EncodeBase64(avatar)} // img := HttpRequest(url)
		account.ImgHash = []uint64{MkImgHash(getImg(avatar))}
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
		Service:   service.Name,
		Username:  username,
		Url:       data.Url,
		Id:        strconv.Itoa(data.Id),
		Bio:       []string{data.Bio},
		Picture:   []string{EncodeBase64(avatar)},
		ImgHash:   []uint64{MkImgHash(getImg(avatar))},
		Location:  data.Location,
		Created:   data.Created_at,
		Updated:   data.Updated_at,
		Blog:      data.Blog,
		Followers: data.Followers,
		Following: data.Following,
	}
	return account
}

func RedditInfo(username string, service Service) Account {
	log.Println("reddit")
	var data struct {
		Id      string `json:"id"`
		Url     string `json:"url"`
		Profile struct {
			Bio       string `json:"bio"`
			Firstname string `json:"firstName"`
			Lastname  string `json:"lastName"`
		} `json:"profile"`
	}
	jsonData := HttpRequest("https://api.reddit.com/user/" + username)
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Println(err)
	}
	return Account{
		Service:   service.Name,
		Username:  username,
		Id:        data.Id,
		Url:       "https://reddit.com/user/" + username,
		Bio:       []string{data.Profile.Bio},
		Firstname: data.Profile.Firstname,
		Lastname:  data.Profile.Lastname,
	}
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
		Bio:       []string{data.Profile.Bio},
		Firstname: data.Profile.Firstname,
		Lastname:  data.Profile.Lastname,
	}
}

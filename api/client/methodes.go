package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/seekr-osint/seekr/api"
)

var (
	ErrNon200StatusCode = errors.New("API returned non-200 status code")
)

func (c *Client) ApiCall(endpoint string) string {
	u := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", c.Ip, c.Port),
		Path:   fmt.Sprintf("/api/%s", strings.Trim(endpoint, "/")),
	}
	return u.String()
}

func (c *Client) GetPerson(id string) (api.Person, error) {
	client := c.HTTPClient
	resp, err := client.Get(c.ApiCall("people/" + id))
	if err != nil {
		return api.Person{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//return api.Person{}, fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)

		return api.Person{}, ErrNon200StatusCode
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return api.Person{}, err
	}

	var person api.Person
	err = json.Unmarshal(body, &person)
	if err != nil {
		return api.Person{}, err
	}

	return person, nil
}
func (c *Client) GetDB() (api.DataBase, error) {
	client := c.HTTPClient
	resp, err := client.Get(c.ApiCall("db"))
	if err != nil {
		return api.DataBase{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//return api.DataBase{}, fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)

		return api.DataBase{}, ErrNon200StatusCode
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return api.DataBase{}, err
	}

	var db api.DataBase
	err = json.Unmarshal(body, &db)
	if err != nil {
		return api.DataBase{}, err
	}

	return db, nil
}

func (c *Client) Ping() (time.Duration, error) {
	startTime := time.Now()

	client := c.HTTPClient
	resp, err := client.Get(c.ApiCall("ping"))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	endTime := time.Now()

	pingTime := endTime.Sub(startTime)
	return pingTime, nil
}

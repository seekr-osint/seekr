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
	"github.com/seekr-osint/seekr/api/functions"
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


func (c *Client) DeletePerson(id string) (error) {
	client := c.HTTPClient
	resp, err := client.Get(c.ApiCall("people/" + id + "/delete"))
	if err != nil {
		return  err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return  ErrNon200StatusCode
	}


	return nil
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
func (c *Client) GetPeople() (map[string]string, error) {
	client := c.HTTPClient
	resp, err := client.Get(c.ApiCall("db"))
	if err != nil {
		return map[string]string{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return map[string]string{}, ErrNon200StatusCode
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]string{}, err
	}

	var db api.DataBase
	err = json.Unmarshal(body, &db)
	if err != nil {
		return map[string]string{}, err
	}
	people := map[string]string{}	
	for _, id := range functions.SortMapKeys(db) {
		people[id] = db[id].Name
	}

	return people, nil
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

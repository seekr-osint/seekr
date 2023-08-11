package client

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

func (c *Client) ApiCall(endpoint string) string {
	u := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", c.Ip, c.Port),
		Path:   fmt.Sprintf("/api/%s", strings.Trim(endpoint, "/")),
	}
	return u.String()
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

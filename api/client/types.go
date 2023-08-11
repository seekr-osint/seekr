package client

import (
	"net/http"
	"time"
)

func NewClient(ip string, port uint32) *Client {
	return &Client{
		Ip: ip,
		Port: uint32(port),
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type Client struct {
	Ip         string
	Port       uint32
	HTTPClient *http.Client
}


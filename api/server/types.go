package server

import "github.com/seekr-osint/seekr/api/webserver"

type ApiServer struct {
	Disable bool `json:"disable"` // Bug unused
}

type Server struct{
	Ip string `json:"ip"`
	Port uint16 `json:"port"`
	WebServer webserver.Webserver `json:"webserver"`
	ApiServer ApiServer `json:"api_server"`
	ForcePort bool `json:"strict_port"` // not changing to the next unused port if port is in use
}

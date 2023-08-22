package webserver

import (
	"embed"
)

type Webserver struct {
	FileSystem embed.FS `json:"file_system"`
	Disable    bool     `json:"disable"`
	LiveServer bool     `json:"live_server"`
	LiveServerPath string `json:"live_server_path"`
}

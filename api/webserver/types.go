package webserver

import (
	"embed"
)

type Webserver struct{
	FileSystem  embed.FS `json:"file_system"`
	Disable bool `json:"disable"`
}

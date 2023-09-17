package config

import "net"

type Config struct {
	Ip           net.IP
	Port         uint16
	DataBasePath string
}

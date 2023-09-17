package config

import "fmt"

func (c Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Ip, c.Port)
}

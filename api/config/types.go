package config

type Server struct {
	Port uint64 `toml:"port" default:"8569"`
	Ip   string `toml:"ip" default:"localhost"`
}
type General struct {
	ForcePort bool `toml:"force_port" default:"false"`
	Browser   bool `toml:"browser" default:"true"`
}
type Config struct {
	Server  Server  `toml:"server"`
	General General `toml:"general"`
}

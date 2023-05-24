package config

type Server struct {
	Port uint64 `toml:"port" example:"8569"`
	Ip   string `toml:"ip" example:"127.0.0.1"`
}
type General struct {
	ForcePort bool `toml:"force_port" example:"false"`
	Browser   bool `toml:"browser" example:"true"`
}
type Config struct {
	Server  Server  `toml:"server"`
	General General `toml:"general"`
}

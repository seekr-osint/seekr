package config

type Server struct {
	Port uint64 `toml:"port" json:"port" default:"8569"`
	Ip   string `toml:"ip" json:"ip" default:"localhost"`
}
type General struct {
	ForcePort bool   `toml:"force_port" json:"force_port" default:"false"`
	Browser   bool   `toml:"browser" json:"browser" default:"true"`
	Discord   bool   `toml:"discord" json:"discord" default:"true"`
	Database  string `toml:"database" json:"-" default:"./data"`
	//CreateConfigFile bool `toml:"create_config_file" default:"false"` // not needed
}
type Config struct {
	Server  Server  `toml:"server" json:"server"`
	General General `toml:"general" json:"general"`
}

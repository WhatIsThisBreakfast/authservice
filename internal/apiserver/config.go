package apiserver

type Config struct {
	Port     string `toml:"server_port"`
	DebugLvl string `toml:"debug_lvl"`
}

func NewConfig() *Config {
	return &Config{
		Port:     ":9999",
		DebugLvl: "debug",
	}
}

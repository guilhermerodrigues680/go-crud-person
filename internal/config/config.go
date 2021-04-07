package config

type Config struct {
	AppName    string
	Version    string
	ServerAddr string
}

func GetConfig() Config {
	return Config{
		AppName:    "go-crud-person",
		Version:    "0.0.1-alpha.0",
		ServerAddr: ":8080",
	}
}

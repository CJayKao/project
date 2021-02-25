package config

type Config struct {
	Port    string
	Address string
}

func InitConfig(path string) Config {
	return Config{
		Port:    "80",
		Address: "127.0.0.1",
	}
}

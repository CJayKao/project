package config

type Config struct {
	Port    string
	Address string
}

func InitConfig() *Config {
	return &Config{
		Port:    "123",
		Address: "address",
	}
}

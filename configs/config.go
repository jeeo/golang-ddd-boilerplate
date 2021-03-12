package configs

import "os"

type DatabaseConfig struct {
	ConnectionString string
	Host             string
	Port             string
	User             string
	Pwd              string
	Name             string
	Static           string
}

type ServerConfig struct {
	Port string
}

type Config struct {
	Server ServerConfig
	Db     DatabaseConfig
}

func ProvideConfig() Config {

	return Config{
		Server: ServerConfig{
			Port: os.Getenv("PORT"),
		},
		Db: DatabaseConfig{
			ConnectionString: os.Getenv("DB_CONNECTION"),
			Host:             os.Getenv("DB_HOST"),
			Port:             os.Getenv("DB_PORT"),
			User:             os.Getenv("DB_USER"),
			Pwd:              os.Getenv("DB_PWD"),
			Name:             os.Getenv("DB_NAME"),
			Static:           "someStaticValue",
		},
	}
}

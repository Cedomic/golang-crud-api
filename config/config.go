package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	SslMode  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Host:     "127.0.0.1",
			Port:     5432,
			Username: "user",
			Password: "password",
			Name:     "musiclibrary",
			SslMode:  "disable",
		},
	}
}

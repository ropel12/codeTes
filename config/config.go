package config

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
}

func NewConfig() Config {
	return Config{
		Username: "test",
		Password: "test",
		Port:     "3306",
		Host:     "188.166.211.2",
		DbName:   "db_test",
	}
}

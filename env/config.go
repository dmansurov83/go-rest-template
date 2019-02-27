package env

type Config struct {
	DbSource string
	DbDriver string
	Port     string
}

func NewConfig() *Config {
	return &Config{
		DbDriver: "sqlite3",
		DbSource: ":memory:",
		Port:     "8000",
	}
}

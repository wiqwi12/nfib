package config

type Config struct {
	Adr    string `env:"ADR_PATH"`
	Token  string `env:"API_TOKEN"`
	DbConn string `env:"DB_CONNECTION"`
}

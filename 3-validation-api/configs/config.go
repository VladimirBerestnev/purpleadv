package configs

type Config struct {
	Email    string
	Password string
	Address  string
}

func NewConfig(e string, password string, address string) *Config {
	return &Config{
		Email:    e,
		Password: password,
		Address:  address,
	}
}

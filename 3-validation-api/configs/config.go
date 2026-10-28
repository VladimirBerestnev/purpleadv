package configs

import "github.com/jordan-wright/email"

type Config struct {
	Email    email.Email
	Password string
	Address  string
}

func GetEmail() *email.Email {
	e := email.NewEmail()
	e.From = "Vladimir <berestnew@ya.ru>"
	e.To = []string{"berestnew@yandex.ru"}
	e.Subject = "Test API"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Test API</h1>")
	return e
}

func NewConfig(e email.Email, password string, address string) *Config {
	return &Config{
		Email:    e,
		Password: password,
		Address:  address,
	}
}

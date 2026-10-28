package verify

import (
	"net/http"
	"net/smtp"
	"validation/configs"
)

type Verify struct{}

func VerifyHandler(router *http.ServeMux) {

	verify := Verify{}

	router.HandleFunc("POST /send", verify.Send())
	router.HandleFunc("GET /verify/{hash}", verify.GetHash)
}

func (verify Verify) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		email := configs.GetEmail()
		config := configs.NewConfig(*email, "password123", "smtp.yandex.ru")

		config.Email.Send("smtp.yandex.ru:587", smtp.PlainAuth("", "test@yandex.ru", config.Password, config.Address))
	}

}

func (verify Verify) GetHash(w http.ResponseWriter, req *http.Request) {

}

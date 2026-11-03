package verify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"validation/configs"

	"github.com/feiin/randomhash"
	"github.com/jordan-wright/email"
)

type Verify struct {
	Email email.Email `json:"email"`
	Hash  string      `json:"hash"`
}

type SendRequest struct {
	Email string `json:"email"`
}

func GetEmail() *email.Email {
	e := email.NewEmail()
	e.From = "Vladimir <test@yandex.ru>"
	e.To = []string{"berestnew@yandex.ru"}
	e.Subject = "Test API"
	e.Text = []byte("Push the link")
	e.HTML = []byte("<h1>Test API</h1>")
	return e
}

func VerifyHandler(router *http.ServeMux) {

	verify := Verify{}

	router.HandleFunc("POST /send", verify.Send())
	router.HandleFunc("GET /verify/{hash}", verify.Validate)
}

func (verify Verify) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		email := GetEmail()
		var SendRequest SendRequest
		json.NewDecoder(req.Body).Decode(&SendRequest)
		email.To = []string{SendRequest.Email}
		hash := hash()
		link := fmt.Sprintf("<a href=\"http://localhost:8082/verify/%s\">http://localhost:8082/verify/%s</a>", hash, hash)
		email.HTML = []byte(link)

		verify.Email = *email
		verify.Hash = hash
		verifyJson, err := json.Marshal(verify)
		if err != nil {
			fmt.Println(err)
		}
		verify.WriteFile(verifyJson)
		config := configs.NewConfig("test@yandex.ru", "password123", "smtp.yandex.ru")

		email.Send("smtp.yandex.ru:587", smtp.PlainAuth("", config.Email, config.Password, config.Address))
	}

}

func hash() string {
	randomHash := randomhash.New("")
	// Генерация хеша длиной 10 символов
	result, err := randomHash.GenerateHash(10)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func (verify Verify) Validate(w http.ResponseWriter, req *http.Request) {
	hash := req.PathValue("hash")
	data, err := ReadFile("data.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return
	}
	var verifyRead Verify
	errJson := json.Unmarshal(data, &verifyRead)
	if errJson != nil {
		fmt.Println("Ошибка Unmarshal")
		return
	}
	if hash != verifyRead.Hash {
		w.Write([]byte("false"))
		return
	}
	os.Remove("data.json")
	w.Write([]byte("true"))
}

func (verify Verify) WriteFile(data []byte) {
	file, errCreate := os.Create("data.json")
	if errCreate != nil {
		fmt.Println("Ошибка создания файла")
		return
	}
	defer file.Close()

	_, errWrite := file.Write(data)
	if errWrite != nil {
		fmt.Println("Ошибка записи в файл")
		return
	}
	fmt.Println("Запись успешна")
}

func ReadFile(filename string) ([]byte, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil

}

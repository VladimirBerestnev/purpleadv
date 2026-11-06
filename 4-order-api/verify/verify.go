package verify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/feiin/randomhash"
)

type Verify struct {
	Hash string
}

func VerifyHandler(router *http.ServeMux) {

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

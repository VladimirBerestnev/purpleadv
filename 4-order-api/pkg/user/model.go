package user

import (
	"fmt"

	"github.com/feiin/randomhash"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Phone     string `json:"phone" validate:"required"`
	SessionID string `json:"sessionId"`
}

type UserRequest struct {
	Phone string `json:"phone" validate:"required"`
}

func NewUser(phone string, sessionId string) *User {
	return &User{
		Phone:     phone,
		SessionID: sessionId,
	}
}

func GenerateSessionID() string {
	result, err := randomhash.New("").GenerateHash(15)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

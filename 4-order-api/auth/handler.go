package auth

import (
	"encoding/json"
	"net/http"
	"order-api/pkg/jwt"
	"order-api/pkg/res"
	"order-api/pkg/user"
	"order-api/req"

	"gorm.io/gorm"
)

func (authHandler *AuthHandler) GetCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[user.UserRequest](&w, r)
		if err != nil {
			return
		}
		sessionId := user.GenerateSessionID()
		userFound, errFound := authHandler.userRepo.GetByPhone(body.Phone)
		if errFound == nil {
			user := user.User{
				Model:     gorm.Model{ID: userFound.ID},
				Phone:     body.Phone,
				SessionID: sessionId,
			}
			updateUser, errUpdate := authHandler.userRepo.Update(&user)
			if errUpdate != nil {
				http.Error(w, errUpdate.Error(), http.StatusBadRequest)
				return
			}
			structIdMarshal, errWrite := WriteSessionID(updateUser.SessionID)
			if errWrite != nil {
				http.Error(w, errWrite.Error(), http.StatusBadRequest)
				return
			}
			w.Write(structIdMarshal)
			return

		}

		user := user.User{Phone: body.Phone, SessionID: sessionId}

		createUser, errCreate := authHandler.userRepo.Create(&user)
		if errCreate != nil {
			http.Error(w, errCreate.Error(), http.StatusBadRequest)
			return
		}
		structIdMarshal, errWrite := WriteSessionID(createUser.SessionID)
		if errWrite != nil {
			http.Error(w, errWrite.Error(), http.StatusBadRequest)
			return
		}
		w.Write(structIdMarshal)
	}
}

func WriteSessionID(sessionId string) ([]byte, error) {
	structId := SessionIdResponse{SessionID: sessionId}
	structIdMarshal, errMarsh := json.Marshal(structId)
	if errMarsh != nil {
		return nil, errMarsh
	}
	return structIdMarshal, nil

}

func (authHandler *AuthHandler) GetToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[AuthCode](&w, r)
		if err != nil {
			return
		}
		user, errGet := authHandler.userRepo.GetBySessionID(body.SessionID)
		if errGet != nil {
			http.Error(w, errGet.Error(), http.StatusNotFound)
			return
		}
		if body.Code != 9999 {
			res.Json(w, "Wrong Code", http.StatusUnauthorized)
			return
		}
		token, errJWT := jwt.NewJWT(authHandler.Config.Auth.Secret).Create(user.Phone)
		if errJWT != nil {
			http.Error(w, errJWT.Error(), http.StatusInternalServerError)
			return
		}
		tokenResponse := jwt.TokenResponse{Token: token}
		tokenMarshal, errMarshal := json.Marshal(tokenResponse)
		if errMarshal != nil {
			http.Error(w, errMarshal.Error(), http.StatusBadRequest)
			return
		}
		w.Write(tokenMarshal)
	}
}

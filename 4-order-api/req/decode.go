package req

import (
	"encoding/json"
	"net/http"
	"order-api/pkg/res"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	var payLoad T
	err := json.NewDecoder(r.Body).Decode(&payLoad)
	if err != nil {
		res.Json(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	errValid := Validate(payLoad)
	if errValid != nil {
		res.Json(*w, errValid.Error(), http.StatusBadRequest)
		return nil, errValid
	}

	return &payLoad, nil

}

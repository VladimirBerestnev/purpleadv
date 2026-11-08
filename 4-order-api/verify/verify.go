package verify

import (
	"encoding/json"
	"net/http"
	"order-api/pkg/product"
	"order-api/pkg/res"

	"github.com/go-playground/validator/v10"
)

func VerifyHandler(router *http.ServeMux) {

	router.HandleFunc("POST /verify", Validate())
}

func Validate() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var product product.Product
		err := json.NewDecoder(req.Body).Decode(&product)
		if err != nil {
			res.Json(w, err.Error(), 402)
			return
		}
		errValid := IsValid(product)
		if errValid != nil {
			res.Json(w, errValid.Error(), 400)
			return
		}
	}

}

func IsValid(product product.Product) error {
	validate := validator.New()
	return validate.Struct(product)
}

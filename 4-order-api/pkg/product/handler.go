package product

import (
	"net/http"
	"order-api/pkg/res"
	"order-api/req"
	"strconv"

	"gorm.io/gorm"
)

func NewProductHandler(router *http.ServeMux, repo ProductHandlerCfg) {
	handler := ProductHandler{
		Repo: repo.Repo,
	}
	router.HandleFunc("POST /product", handler.Create())
	router.HandleFunc("GET /product/{id}", handler.GetById())
	router.HandleFunc("PATCH /product/{id}", handler.Update())
	router.HandleFunc("DELETE /product/{id}", handler.Delete())
}

func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductRequest](&w, r)
		if err != nil {
			return
		}
		product := Product{Name: body.Name, Desc: body.Desc, Price: body.Price, Images: body.Images}

		createProduct, errCreate := handler.Repo.Create(&product)
		if errCreate != nil {
			http.Error(w, errCreate.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createProduct, http.StatusCreated)
	}
}

func (handler *ProductHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		product, err := handler.Repo.GetById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		res.Json(w, product, http.StatusOK)
	}
}

func (handler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductUpdate](&w, r)
		if err != nil {
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		product, err := handler.Repo.Update(&Product{
			Model:  gorm.Model{ID: uint(id)},
			Name:   body.Name,
			Desc:   body.Desc,
			Price:  body.Price,
			Images: body.Images,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, product, http.StatusOK)
	}
}

func (handler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = handler.Repo.GetById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		errDelete := handler.Repo.Delete(id)
		if errDelete != nil {
			http.Error(w, errDelete.Error(), http.StatusBadRequest)
			return
		}
	}
}

package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

var validate = validator.New()

func ValidateStruct(r *http.Request, w http.ResponseWriter, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	if err := validate.Struct(v); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}
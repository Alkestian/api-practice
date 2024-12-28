package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Alkestian/api-practice/utils"
	"github.com/go-chi/chi"
)

type User struct{
	ID string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

var hardCodedUsers []User

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser")
	var user User
	if err := utils.ValidateStruct(r, w, &user); err != nil {
		return
	}
	userJSON, err := json.Marshal(user); if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hardCodedUsers = append(hardCodedUsers, user)
	w.WriteHeader(http.StatusCreated) 
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}

func (u *User) ListUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ListUsers")
	usersJSON, err := json.Marshal(hardCodedUsers); if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(usersJSON)
}

func (u *User) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	for _, user := range hardCodedUsers {
		if user.ID == id {
			userJSON, err := json.Marshal(user); if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(userJSON)
			return
		}
	}
}

func (u *User) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateUser")
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	if err := utils.ValidateStruct(r, w, &User{}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, user := range hardCodedUsers {
		if user.ID == id {
			var newUser User
			if err := utils.ValidateStruct(r, w, &newUser); err != nil {
				return
			}
			hardCodedUsers[i] = newUser
			break
		}
	}
	w.WriteHeader(http.StatusOK)
}

func (u *User) UpdateUserField(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("UpdateUserField")
}

func (u *User) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteUser")
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	for i, user := range hardCodedUsers {
		if user.ID == id {
			hardCodedUsers = append(hardCodedUsers[:i], hardCodedUsers[i+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusOK)
}
package handler

import (
	"fmt"
	"net/http"
)

type User struct{}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Println("CreateUser")
}

func (u *User) ListUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("ListUsers")
}

func (u *User) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("GetUser")
}

func (u *User) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("UpdateUser")
}

func (u *User) UpdateUserField(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("UpdateUserField")
}

func (u *User) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("DeleteUser")
}
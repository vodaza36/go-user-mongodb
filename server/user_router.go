package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/vodaza36/go-user-mongodb/pck"

	"github.com/gorilla/mux"
)

type userRouter struct {
	userService root.UserService
}

// NewUserRouter create a router instance
func NewUserRouter(u root.UserService, router *mux.Router) *mux.Router {
	userRouter := userRouter{u}

	router.HandleFunc("/", userRouter.createUserHandler).Methods("PUT")
	router.HandleFunc("/{username}", userRouter.getUserHandler).Methods("GET")
	return router
}

func (ur *userRouter) createUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := decodeUser(r)
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = ur.userService.CreateUser(&user)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	JSON(w, http.StatusOK, err)
}

func (ur *userRouter) getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	username := vars["username"]

	user, err := ur.userService.GetByUsername(username)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	JSON(w, http.StatusOK, user)
}

func decodeUser(ur *http.Request) (root.User, error) {
	var u root.User
	if ur.Body == nil {
		return u, errors.New("no request body")
	}
	decoder := json.NewDecoder(ur.Body)
	err := decoder.Decode(&u)
	return u, err
}

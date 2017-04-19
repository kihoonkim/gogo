package user

import (
	"github.com/facebookgo/inject"
	"github.com/gorilla/mux"
	"github.com/kihoonkim/gogo/utils"
	"net/http"
)

type Handler struct {
	Service *Service `inject:""`
}

var u Handler

func init() {
	inject.Populate(&u)
}

func HandlerMapping(router *mux.Router) {
	router.HandleFunc("/users", u.createUserHandler).Methods("POST")
	router.HandleFunc("/users", u.findAllUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}", u.findUserHandler).Methods("GET")
}

func (u *Handler) createUserHandler(resp http.ResponseWriter, req *http.Request) {
	var user User
	utils.DecodeRequestBody(resp, req, &user)

	u.Service.Save(&user)

	utils.ResponseWithJson(resp, http.StatusOK, user)
}

func (u *Handler) findAllUserHandler(resp http.ResponseWriter, req *http.Request) {
	users := u.Service.FindAll()
	utils.ResponseWithJson(resp, http.StatusOK, users)
}

func (u *Handler) findUserHandler(resp http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["id"]

	user := u.Service.FindById(userId)

	utils.ResponseWithJson(resp, http.StatusOK, user)
}

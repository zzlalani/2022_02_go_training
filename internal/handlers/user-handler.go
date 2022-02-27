package handlers

import (
	"github.com/gorilla/mux"
	"github.com/zzlalani/go-practice/internal/dto"
	"github.com/zzlalani/go-practice/internal/services"
	"net/http"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler (userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService,
	}
}

func (h *UserHandler) Setup(router *mux.Router) {
	router.HandleFunc("/", h.PostUsers).Methods("POST")
	//r.HandleFunc("/", GetUsers).Methods("GET")
	//r.HandleFunc("/{id}", GetUserByID).Methods("GET")
	//r.HandleFunc("/{id}", UpdateUser).Methods("PUT")
	//r.HandleFunc("/{id}", DeleteUser).Methods("DELETE")
}

func (h *UserHandler) PostUsers(w http.ResponseWriter, r *http.Request) {
	body := dto.UserReq{}
	Bind(r.Body, &body)
	id, err := h.userService.CreateUser(body)
	if err != nil {
		Error(w, "unable to create user", http.StatusInternalServerError)
	}
	resp := dto.UserRes{
		ID: id,
	}
	Response(w, resp, http.StatusOK)
}

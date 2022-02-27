package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zzlalani/go-practice/internal/repository"
	"github.com/zzlalani/go-practice/internal/services"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Setup(db *gorm.DB) {
	r := mux.NewRouter()
	// user routes
	userService := services.NewUserService(repository.NewUserRepo(db))
	usersRouter := r.PathPrefix("/users").Subrouter()
	userHandler := NewUserHandler(userService)
	userHandler.Setup(usersRouter)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

func Error(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf("{message: %s}", message)))
}

func Response(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp, err := UnBind(data)
	if err != nil {
		panic(err)
	}
	w.Write(resp)
}

// Bind json string to dto
func Bind(r io.Reader, v interface{}) error {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, v)
	if err != nil {
		return err
	}
	return nil
}

func UnBind(v interface{}) ([]byte, error) {
	resp, err := json.Marshal(v)
	if err != nil {
		return []byte(""), err
	}
	return resp, nil
}

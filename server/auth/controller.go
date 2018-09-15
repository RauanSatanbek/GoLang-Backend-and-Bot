package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"makebex-backend/server/config"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

var DB *sql.DB
var (
	basePath = "/auth"
	getPath  = "/auth/{id}"

)

func Initialize(router *mux.Router, db *sql.DB) {
	DB = db

	Migration(DB)
	DefineRouter(router)
}

func DefineRouter(router *mux.Router) {
	//-* GET
	router.
		HandleFunc(getPath, GetUser).
		Methods(config.Get)

	//-* POST
	router.
		HandleFunc(basePath, CreateUser).
		Methods(config.Post)

	fmt.Println("Auth: router - OK")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi( vars["id"] )

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(ErrorResponse{
			Status: http.StatusBadRequest,
			Message: "user_id is not valid or empty",
		})
		return
	}

	user := User{ ID: userId }
	user.Get(DB)

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := User{}

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(ErrorResponse{
			Status: http.StatusBadRequest,
			Message: "Payload not valid.",
		})
		return
	}

	user.Create(DB)

	json.NewEncoder(w).Encode(user)
}


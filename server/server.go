package server

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"makebex-backend/server/auth"
	"makebex-backend/server/config"
	"net/http"
)

type Server struct {
	Router *mux.Router
	DB *sql.DB
}

func On(addr string, db *sql.DB) {
	s := Server{ DB:db }

	s.useRouters()
	s.Run(addr)
}

func (s *Server) useRouters() {
	/*- /api/v1 -*/
	s.Router = mux.NewRouter().
		PathPrefix(config.VersionOne).
		Subrouter()

	auth.
		Initialize(s.Router, s.DB)
}

func (s *Server) Run(addr string) {
	fmt.Printf("Listening on - %s", addr)

	err := http.ListenAndServe(addr, s.Router)
	if err != nil {
		log.Fatal(err)
	}
}
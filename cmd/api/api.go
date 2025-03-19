package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/bishalkl/learnBackendApi/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIserver(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {

	// let's create router now
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// user handler
	userHandler := user.NewHandler()
	userHandler.RegisterRouter(subrouter)

	//  login and server listen
	log.Println("Listening on ", s.addr)
	return http.ListenAndServe(s.addr, router)
}

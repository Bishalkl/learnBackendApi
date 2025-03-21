package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/bishalkl/learnBackendApi/service/auth"
	"github.com/bishalkl/learnBackendApi/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	db    *sql.DB
	store *user.Store //Store for interacting with the database
}

func NewAPIserver(addr string, db *sql.DB) *APIServer {

	store := user.NewStore(db)
	return &APIServer{
		addr:  addr,
		db:    db,
		store: store,
	}
}

func (s *APIServer) Run() error {

	// let's create router now
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// user handler
	userHandler := user.NewHandler(s.store)
	userHandler.RegisterRouter(subrouter)

	// Apply middleware to the subrouter (the route under /api/v1)
	subrouter.Use(auth.LoggingMiddleware)
	subrouter.Use(auth.JWTMiddleware)

	//  login and server listen
	log.Println("Listening on ", s.addr)
	return http.ListenAndServe(s.addr, router)
}

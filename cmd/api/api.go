package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/bishalkl/learnBackendApi/service/auth"
	"github.com/bishalkl/learnBackendApi/service/product"
	"github.com/bishalkl/learnBackendApi/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr         string
	db           *sql.DB
	userstore    *user.Store    //Store for interacting with the database
	productstore *product.Store //Store for interacting with the database
}

func NewAPIserver(addr string, db *sql.DB) *APIServer {

	userstore := user.NewStore(db)
	productstore := product.NewStore(db)
	return &APIServer{
		addr:         addr,
		db:           db,
		userstore:    userstore,
		productstore: productstore,
	}
}

func (s *APIServer) Run() error {

	// let's create router now
	router := mux.NewRouter()
	// for public router
	publirouter := router.PathPrefix("/api/v1").Subrouter()
	// for protectedrouter
	protectedrouter := router.PathPrefix("/protected/api/v1").Subrouter()

	// user handler
	userHandler := user.NewHandler(s.userstore)
	userHandler.RegisterRouter(publirouter)

	// Product handler
	productHanlder := product.NewHandler(s.productstore)
	productHanlder.RegisterRouter(publirouter)

	// Apply middleware to the subrouter (the route under /api/v1)
	publirouter.Use(auth.LoggingMiddleware)

	// Apply middleware to the subrouter(the router under /protected/api/v1)
	protectedrouter.Use(auth.JWTMiddleware, auth.LoggingMiddleware)

	//  login and server listen
	log.Println("Listening on ", s.addr)
	return http.ListenAndServe(s.addr, router)
}

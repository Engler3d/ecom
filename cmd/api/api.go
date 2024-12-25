package api

// creating new struct for server interface
// creating construnction function for assinging values
// creating Run method for starting server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Engler3d/ecom/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string  // adress
	db   *sql.DB // database connection
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegistearRoutes(subrouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

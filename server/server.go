package server

import (
	"log"
	"net/http"
	"os"

	"github.com/vodaza36/go-user-mongodb/pck"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server class
type Server struct {
	router *mux.Router
}

// NewServer creates a Server instance
func NewServer(u root.UserService) *Server {
	s := Server{router: mux.NewRouter()}
	NewUserRouter(u, s.newSubrouter("/user"))
	return &s
}

// Start the server to listen
func (s *Server) Start() {
	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, s.router)); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func (s *Server) newSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}

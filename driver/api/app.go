package api

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/devbenatti/rbwallet-go/config"
	action "github.com/devbenatti/rbwallet-go/driver/api/action/account"
	"github.com/gorilla/mux"
)

type Server struct {
	config config.Config
	router *mux.Router
}

func NewServer() *Server {
	return &Server{
		config: config.Load(),
		router: mux.NewRouter(),
	}
}

func (s *Server) configureHandlers() {
	createAccount := action.NewCreateAccount()

	s.router.HandleFunc("/account", createAccount.Execute()).Methods(http.MethodPost)
}

func (s *Server) Listen() {
	s.configureHandlers()

	http.Handle("/", s.router)

	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":8080",
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

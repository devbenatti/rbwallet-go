package api

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/devbenatti/rbwallet-go/config"
	accountAction "github.com/devbenatti/rbwallet-go/driver/api/action/account"
	transactionAction "github.com/devbenatti/rbwallet-go/driver/api/action/transaction"
	"github.com/devbenatti/rbwallet-go/driver/api/middleware"
	"github.com/urfave/negroni"

	"github.com/gorilla/mux"
)

type Server struct {
	config config.Config
	router *mux.Router
	ngn    *negroni.Negroni
}

func NewServer() *Server {
	return &Server{
		config: config.Load(),
		router: mux.NewRouter(),
		ngn:    negroni.New(),
	}
}

func (s *Server) configureHandlers() {
	createAccount := accountAction.NewCreateAccount()
	findAccount := accountAction.NewFindAccount()
	payloadValidator := middleware.NewPayloadValidator()

	s.router.Handle("/accounts", s.ngn.With(
		negroni.Wrap(payloadValidator.Execute()),
		negroni.Wrap(createAccount.Execute()),
	)).Methods(http.MethodPost)

	s.router.Handle("/accounts/{id}", findAccount.Execute()).Methods(http.MethodGet)

	createTransaction := transactionAction.NewCreateTransaction()

	s.router.Handle("/transactions", createTransaction.Execute()).Methods(http.MethodPost)
}

func (s *Server) Listen() {
	s.configureHandlers()

	http.Handle("/", s.router)

	s.ngn.Use(negroni.NewLogger())
	s.ngn.Use(negroni.NewRecovery())

	s.ngn.UseHandler(s.router)

	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":8080",
		Handler:      s.ngn,
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

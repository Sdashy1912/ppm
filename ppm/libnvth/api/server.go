package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/spf13/viper"
)

// Server provides an http.Server.
type Server struct {
	*http.Server
}

// NewServer creates and configures an APIServer serving all application routes.
func NewServer() (*Server, error) {
	log.Println("configuring server...")
	// api, err := New(viper.GetBool("enable_cors"))
	api, err := New()
	if err != nil {
		return nil, err
	}

	var addr string
	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	if host == "" || port == "" {
		return nil, errors.New("Could not load server configurations")
	}
	// allow port to be set as localhost:3000 in env during development to avoid "accept incoming network connection" request on restarts
	addr = fmt.Sprintf("%s:%s", host, port)

	srv := http.Server{
		Addr:    addr,
		Handler: api,
	}

	return &Server{&srv}, nil
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (srv *Server) Start() {
	log.Println("starting server...")
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("Listening on %s\n", srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Println("Shutting down server... Reason:", sig)
	// teardown logic...

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Println("Server gracefully stopped")
}

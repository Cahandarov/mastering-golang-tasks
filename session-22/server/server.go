package server

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"session-22/handler"
)

func Init() error {
	log.Info("Starting server at port: 8080")
	return (&http.Server{
		Handler: handler.GetRouter(),
		Addr:    ":8080",
	}).ListenAndServe()
}

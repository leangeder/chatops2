package server

import (
	"github.com/gorilla/mux"
	"github.com/leangeder/chatops2/pkg/features"
	"github.com/leangeder/chatops2/pkg/server/logger"
	"github.com/leangeder/chatops2/pkg/server/metrics"
)

type server struct {
	Router *mux.Router
}

func Run() (*server, error) {
	s := &server{}
	s.Router = mux.NewRouter().StrictSlash(true)
	s.Router.Use(logger.Logger)
	s.Router.HandleFunc("/metrics", metrics.Metrics).Methods("GET")
	features.LoadFeatures(s.Router)

	return s, nil
}

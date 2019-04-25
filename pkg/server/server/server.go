package server

import (
	"github.com/gorilla/mux"
	"github.com/leangeder/chatops2/pkg/features"
	"github.com/leangeder/chatops2/pkg/server/logger"
)

type server struct {
	Router *mux.Router
}

func Run() (*server, error) {
	s := &server{}
	s.Router = mux.NewRouter().StrictSlash(true)
	s.Router.Use(logger.Logger)
	features.LoadFeatures(s.Router)

	return s, nil
}

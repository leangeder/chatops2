package server

import (
	"github.com/gorilla/mux"
	"github.com/leangeder/chatops2/pkg/features"
)

type Server struct {
	Handlers *mux.Router
}

func Run() *Server {
	s := new(Server)

	s.Handlers = features.NewRouters()
	// s.Handlers = router.NewRouterV1()

	return s
}

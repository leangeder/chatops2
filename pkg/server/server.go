package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leangeder/chatops2/pkg/server/build"
	"github.com/leangeder/chatops2/pkg/server/router"
	"github.com/leangeder/chatops2/pkg/slack"
)

type Server struct {
	Router   http.Handler
	Handlers *mux.Router
	// Projects     map[string]project.Project
	Builds       chan build.Build
	Interactions chan slack.SlackInteraction
}

func Run() *Server {
	s := new(Server)

	s.Builds = make(chan build.Build, 5)
	s.Interactions = make(chan slack.SlackInteraction, 5)
	// s.Projects = make(map[string]project.Project)
	s.Handlers = router.NewRouterV1()

	return s
}

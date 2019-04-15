package pipeline

import (
	"github.com/gorilla/mux"
	"github.com/leangeder/chatops2/pkg/server/slack"
)

type pipeline struct {
	Projects     map[string]Project
	Builds       chan Build
	Interactions chan slack.SlackInteraction
}

type Build struct {
	Project Project
	Target  string
	Image   string
	Type    string
}

func Start(r *mux.Router) {
	p := &pipeline{}
	p.Builds = make(chan Build, 5)
	p.Interactions = make(chan slack.SlackInteraction, 5)
	p.Projects = make(map[string]Project)

	p.Processors()
	p.LoadProject()

	r.HandleFunc("/build-complete", p.BuildComplete).Name("BuildComplete").Methods("POST")
	r.HandleFunc("/slack-interactions", p.SlackInteractions).Name("SlackInteractions").Methods("POST")
}

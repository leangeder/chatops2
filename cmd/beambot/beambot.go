package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/leangeder/chatops2/pkg/server/server"
)

func main() {
	// s := &server.Server{}

	// s.Builds = make(chan build.Build, 5)
	// s.Interactions = make(chan slack.SlackInteraction, 5)
	// // s.Projects = make(map[string]project.Project)
	// s.Handlers = router.NewRouterV1()
	s := server.Run()

	// create router and start listen on port 8000
	fmt.Println("listening on port 8080")
	// log.Fatal(http.ListenAndServe(":8080", middleware.SetupGlobalMiddleware(s.Handlers)))
	log.Fatal(http.ListenAndServe(":8080", s.Handlers))
}

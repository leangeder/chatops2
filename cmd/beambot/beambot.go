package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/leangeder/chatops2/pkg/server"
	"github.com/leangeder/chatops2/pkg/server/middleware"
	"github.com/spf13/viper"
)

func setupConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	setupConfig()
}

func main() {
	// s := &server.Server{}

	// s.Builds = make(chan build.Build, 5)
	// s.Interactions = make(chan slack.SlackInteraction, 5)
	// // s.Projects = make(map[string]project.Project)
	// s.Handlers = router.NewRouterV1()
	s := server.Run()

	// create router and start listen on port 8000
	fmt.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", middleware.SetupGlobalMiddleware(s.Handlers)))
}

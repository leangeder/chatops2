package features

import (
	"github.com/gorilla/mux"
	"github.com/leangeder/chatops2/pkg/features/kubernetes"
	"github.com/leangeder/chatops2/pkg/features/pipeline"
	"github.com/leangeder/chatops2/pkg/server/logger"
)

func NewRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(logger.Logger)
	sub := router.PathPrefix("/v1").Subrouter()
	pipeline.Start(sub)
	kubernetes.Start(sub)
	return router
}

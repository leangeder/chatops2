package features

import (
	"github.com/gorilla/mux"
	"github.com/leangeder/chatops2/pkg/features/kubernetes"
	"github.com/leangeder/chatops2/pkg/features/pipeline"
)

type Features struct {
	Router *mux.Router
	test   map[string]interface{}
}

func LoadFeatures(r *mux.Router) error {
	s := &Features{Router: r.PathPrefix("/v1").Subrouter()}
	s.loadHandlers()
	return nil
}

func (f *Features) loadHandlers() error {
	pipeline.Start(f.Router)
	kubernetes.Start(f.Router)
	return nil
}

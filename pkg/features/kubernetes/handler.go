package kubernetes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Start(r *mux.Router) {
	r.HandleFunc("/test", GetTest).Name("test").Methods("GET")
}

func GetTest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hi dddfg there, I love %s!", r.URL.Path[1:])
}

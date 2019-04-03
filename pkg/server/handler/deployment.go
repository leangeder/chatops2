package handler

import (
	"fmt"
	"net/http"
)

// GetLogs is an httpHandler for route GET /k8s/logs/{applicationName}
func DeploymentToPreview(w http.ResponseWriter, r *http.Request) {
	// // // swagger:route GET /k8s/logs/{applicationName} log logPods
	// // //
	// // // Lists all people.
	// // //
	// // // This will show all recorded people.
	// // //
	// // //     Consumes:
	// // //     - application/json
	// // //
	// // //     Produces:
	// // //     - application/json
	// // //
	// // //     Schemes: http, https
	// // //
	// // //     Responses:
	// // //       200: peopleResponse
	// // params := mux.Vars(r)
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// kubernetes.DeploymentToPreview()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	// fmt.Fprintf(w, kubernetes.GetLogs())
	//for _, item := range kubernetes.GetLogs(params["name"]) {
	//	if item.ID == params["id"] {
	//		w.WriteHeader(http.StatusOK)
	//		// add a arbitraty pause of 1 second
	//		time.Sleep(1000 * time.Millisecond)
	//		if err := json.NewEncoder(w).Encode(item); err != nil {
	//			panic(err)
	//		}
	//		return
	//	}
	//}
	//// If we didn't find it, 404
	//w.WriteHeader(http.StatusNotFound)
	//if err := json.NewEncoder(w).Encode(jsonError{Message: "Not Found"}); err != nil {
	//	panic(err)
	//}
}

func GetTest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

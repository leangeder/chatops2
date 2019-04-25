package pipeline

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func (p *pipeline) BuildComplete(w http.ResponseWriter, r *http.Request) {

	projectName := r.FormValue("project")
	project, ok := p.Projects[projectName]
	if !ok {
		log.Println(errors.New("Project " + projectName + " not found"))
		http.Error(w, "Error encountered", 500)
		return
	}

	build := Build{}
	build.Project = project
	build.Image = r.FormValue("image")   //docker image
	build.Target = r.FormValue("target") // name of the branch or tag
	build.Type = r.FormValue("type")     // branch or tag

	p.Builds <- build
	w.Write([]byte("Received successfully"))
}

func (p *pipeline) SlackInteractions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hi dddfg there, I love %s!", r.URL.Path[1:])
}

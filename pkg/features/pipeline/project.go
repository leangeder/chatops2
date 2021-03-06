package pipeline

type Project struct {
	Name    string
	ID      string
	URL     string
	Channel string
	QA      []string
	Owners  []string
}

func (p *pipeline) LoadProject() {
	var projects []Project
	projects = append(projects, Project{
		Name:    "api-core",
		ID:      "beamery-trials",
		URL:     "https://13.34.34.25",
		Channel: "#general",
		// Channel: "https://leangeder.slack.com/messages/CD8KHAKA8",
		// // Channel: "slack/dsd/deploy-trials",
		QA:     []string{"erleaine", "bart"},
		Owners: []string{"gregory", "erleaine"},
	})

	for _, _project := range projects {
		p.Projects[_project.ID] = _project
	}
}

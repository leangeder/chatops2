package pipeline

type Project struct {
	ID      string
	URL     string
	Channel string
	QA      []string
	Owners  []string
}

func (p *pipeline) LoadProject() {
	var projects []Project
	projects = append(projects, Project{
		ID:      "beamery-trials",
		URL:     "https://13.34.34.25",
		Channel: "slack/dsd/deploy-trials",
		QA:      []string{"erleaine", "bart"},
		Owners:  []string{"gregory", "erleaine"},
	})

	for _, _project := range projects {
		p.Projects[_project.ID] = _project
	}
}

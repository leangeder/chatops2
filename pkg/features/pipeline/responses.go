package pipeline

import (
	"encoding/json"

	"github.com/leangeder/chatops2/pkg/server/slack"
)

func sendSuccessProdDeploy(payload actionPayload, user, url string) (err error) {

	project := payload.Build.Project

	var newM slack.SlackMessage
	newM.Channel = project.Channel
	newM.Text = "New production deployment for project " + project.Name + " by <@" + user + ">"

	newM.Attachments = []slack.SlackAttachment{
		slack.SlackAttachment{
			Fallback: "Project: " + project.Name + " Image: " + payload.Build.Image + " By: <@" + user + ">",
			Fields: []slack.SlackField{
				slack.SlackField{
					Title: "Project",
					Value: project.Name,
					Short: false,
				},
				slack.SlackField{
					Title: "Docker Image",
					Value: payload.Build.Image,
					Short: false,
				},
				slack.SlackField{
					Title: "By",
					Value: "<@" + user + ">",
					Short: true,
				},
			},
		},
		slack.SlackAttachment{
			Fallback: "View project: " + url,
			Color:    "good",
			Actions: []slack.SlackAction{
				slack.SlackAction{
					Type: "button",
					Text: "View project",
					URL:  url,
				},
			},
		},
	}

	_, err = slack.SendSlack(newM)
	return
}

func sendFailedProdDeploy(payload actionPayload, deployErr error) (err error) {

	project := payload.Build.Project

	var newM slack.SlackMessage
	newM.Channel = project.Channel
	newM.Text = "Production deployment failed for project " + project.Name

	newM.Attachments = []slack.SlackAttachment{
		slack.SlackAttachment{
			Fallback: "Project: " + project.Name + " Image: " + payload.Build.Image,
			Fields: []slack.SlackField{
				slack.SlackField{
					Title: "Project",
					Value: project.Name,
					Short: false,
				},
				slack.SlackField{
					Title: "Docker Image",
					Value: payload.Build.Image,
					Short: false,
				},
			},
		},
		slack.SlackAttachment{
			Fallback: "Failure Reason: " + deployErr.Error(),
			Color:    "danger",
			Fields: []slack.SlackField{
				slack.SlackField{
					Title: "Failure Reason",
					Value: deployErr.Error(),
					Short: false,
				},
			},
		},
	}

	_, err = slack.SendSlack(newM)
	return
}

func sendAttemptDeployMessage(build Build) (ts string, err error) {
	msg := getAttemptDeployMessage(build)

	resp, err := slack.SendSlack(msg)
	if err != nil {
		return
	}

	attemptMsgResp := make(map[string]string)
	json.Unmarshal(resp, &attemptMsgResp)
	ts = attemptMsgResp["ts"]

	return
}

func getAttemptDeployMessage(build Build) slack.SlackMessage {
	return slack.SlackMessage{
		Channel: build.Project.Channel,
		Text:    "New Build complete.\nAttempting deployment...",
		Attachments: []slack.SlackAttachment{
			slack.SlackAttachment{
				Fallback: "Project: " + build.Project.Name + " Type: " + build.Type + " Target: " + build.Target + " Image: " + build.Image,
				Fields: []slack.SlackField{
					slack.SlackField{
						Title: "Project",
						Value: build.Project.Name,
						Short: false,
					},
					slack.SlackField{
						Title: "Docker Image",
						Value: build.Image,
						Short: false,
					},
					slack.SlackField{
						Title: "Type",
						Value: build.Type,
						Short: true,
					},
					slack.SlackField{
						Title: "Target",
						Value: build.Target,
						Short: true,
					},
				},
			},
		},
	}
}

func sendDeploySuccessMessage(build Build, ts, url string) (err error) {
	msg := getDeploySuccessMessage(build, url)
	msg.Update = true
	msg.Ts = ts

	_, err = slack.SendSlack(msg)
	return
}

func getDeploySuccessMessage(build Build, url string) slack.SlackMessage {
	return slack.SlackMessage{
		Channel: build.Project.Channel,
		Text:    "New Build complete.\nDeployment Successful! :sunglasses:",
		Attachments: []slack.SlackAttachment{
			slack.SlackAttachment{
				Fallback: "Project: " + build.Project.Name + " Type: " + build.Type + " Target: " + build.Target + " Image: " + build.Image,
				Fields: []slack.SlackField{
					slack.SlackField{
						Title: "Project",
						Value: build.Project.Name,
						Short: false,
					},
					slack.SlackField{
						Title: "Docker Image",
						Value: build.Image,
						Short: false,
					},
					slack.SlackField{
						Title: "Type",
						Value: build.Type,
						Short: true,
					},
					slack.SlackField{
						Title: "Target",
						Value: build.Target,
						Short: true,
					},
				},
			},
			slack.SlackAttachment{
				Fallback: "View project: " + url,
				Color:    "good",
				Actions: []slack.SlackAction{
					slack.SlackAction{
						Type: "button",
						Text: "View project",
						URL:  url,
					},
				},
			},
		},
	}
}

func sendFailedDeployMessage(build Build, ts string, deployErr error) (err error) {
	msg := getFailedDeployMessage(build, deployErr)
	msg.Update = true
	msg.Ts = ts

	_, err = slack.SendSlack(msg)
	return
}

func getFailedDeployMessage(build Build, err error) slack.SlackMessage {
	return slack.SlackMessage{
		Channel: build.Project.Channel,
		Text:    "New Build complete.\nDeployment Failed :sob:",
		Attachments: []slack.SlackAttachment{
			slack.SlackAttachment{
				Fallback: "Project: " + build.Project.Name + " Type: " + build.Type + " Target: " + build.Target + " Image: " + build.Image,
				Fields: []slack.SlackField{
					slack.SlackField{
						Title: "Project",
						Value: build.Project.Name,
						Short: false,
					},
					slack.SlackField{
						Title: "Docker Image",
						Value: build.Image,
						Short: false,
					},
					slack.SlackField{
						Title: "Type",
						Value: build.Type,
						Short: true,
					},
					slack.SlackField{
						Title: "Target",
						Value: build.Target,
						Short: true,
					},
				},
			},
			slack.SlackAttachment{
				Fallback: "Failure Reason: " + err.Error(),
				Color:    "danger",
				Fields: []slack.SlackField{
					slack.SlackField{
						Title: "Failure Reason",
						Value: err.Error(),
						Short: false,
					},
				},
			},
		},
	}
}

func sendOwnerMessages(build Build, url string) (
	payload actionPayload, errs []error) {
	var oMsgs []ownerMsg

	successMessage := getDeploySuccessMessage(build, url)

	for _, user := range build.Project.Owners {
		successMessage.Channel = user
		resp, err := slack.SendSlack(successMessage)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		respMap := make(map[string]string)
		json.Unmarshal(resp, &respMap)

		oMsgs = append(oMsgs, ownerMsg{
			Owner:   user,
			Ts:      respMap["ts"],
			Channel: respMap["channel"],
		})
	}

	payload = actionPayload{
		Build:         build,
		OwnerMessages: oMsgs,
	}

	OwnerMessage, err := getOwnerMessage(build, url, payload)
	if err != nil {
		errs = append(errs, err)
		return
	}

	for _, oM := range payload.OwnerMessages {
		OwnerMessage.Update = true
		OwnerMessage.Channel = oM.Channel
		OwnerMessage.Ts = oM.Ts

		_, err := slack.SendSlack(OwnerMessage)
		if err != nil {
			errs = append(errs, err)
			continue
		}
	}

	return
}

func getOwnerMessage(build Build, url string, payload actionPayload) (slack.SlackMessage, error) {

	marshaledPayload, err := json.Marshal(payload)
	if err != nil {
		return slack.SlackMessage{}, err
	}

	qaTeamAttachment := getQaSlackAttachment(build)

	message := slack.SlackMessage{
		Text: "New Build complete.\nDeployment Successful! :sunglasses:",
		Attachments: []slack.SlackAttachment{
			slack.SlackAttachment{
				Fallback: "Project: " + build.Project.Name + " Type: " + build.Type + " Target: " + build.Target + " Image: " + build.Image,
				Fields: []slack.SlackField{
					slack.SlackField{
						Title: "Project",
						Value: build.Project.Name,
						Short: false,
					},
					slack.SlackField{
						Title: "Docker Image",
						Value: build.Image,
						Short: false,
					},
					slack.SlackField{
						Title: "Type",
						Value: build.Type,
						Short: true,
					},
					slack.SlackField{
						Title: "Target",
						Value: build.Target,
						Short: true,
					},
				},
			},
			slack.SlackAttachment{
				Fallback: "View project: " + url,
				Color:    "good",
				Actions: []slack.SlackAction{
					slack.SlackAction{
						Type: "button",
						Text: "View project",
						URL:  url,
					},
				},
			},
			qaTeamAttachment,
			slack.SlackAttachment{
				Fallback:   "Deploy to Production.",
				CallbackID: "Deploy Decision",
				Actions: []slack.SlackAction{
					slack.SlackAction{
						Type:  "button",
						Text:  "Deploy to Production",
						Name:  "deploy",
						Value: string(marshaledPayload),
						Style: "primary",
						Confirm: map[string]string{
							"title":        "Are you sure?",
							"text":         "This will deploy to production. The process cannot be reversed.",
							"ok_text":      "Deploy",
							"dismiss_text": "Cancel",
						},
					},
					slack.SlackAction{
						Type:  "button",
						Text:  "Close",
						Name:  "close",
						Value: string(marshaledPayload),
						Style: "danger",
						Confirm: map[string]string{
							"title":        "Are you sure?",
							"text":         "This will close this deployment. The process cannot be reversed.",
							"ok_text":      "Close",
							"dismiss_text": "Cancel",
						},
					},
				},
			},
		},
	}

	return message, nil
}

func sendQaMessages(build Build, url string, payload actionPayload) (errs []error) {

	QAmsg, err := getQAMessage(build, url, payload)
	if err != nil {
		errs = append(errs, err)
		return
	}

	for _, user := range build.Project.QA {
		QAmsg.Channel = user
		_, err := slack.SendSlack(QAmsg)
		if err != nil {
			errs = append(errs, err)
			continue
		}
	}

	return
}

func getQAMessage(build Build, url string, payload actionPayload) (slack.SlackMessage, error) {

	marshaledPayload, err := json.Marshal(payload)
	if err != nil {
		return slack.SlackMessage{}, err
	}

	message := slack.SlackMessage{
		Text: "New Build complete.\nDeployment Successful! :sunglasses:",
		Attachments: []slack.SlackAttachment{
			slack.SlackAttachment{
				Fallback: "Project: " + build.Project.Name + " Type: " + build.Type + " Target: " + build.Target + " Image: " + build.Image,
				Fields: []slack.SlackField{
					slack.SlackField{
						Title: "Project",
						Value: build.Project.Name,
						Short: false,
					},
					slack.SlackField{
						Title: "Docker Image",
						Value: build.Image,
						Short: false,
					},
					slack.SlackField{
						Title: "Type",
						Value: build.Type,
						Short: true,
					},
					slack.SlackField{
						Title: "Target",
						Value: build.Target,
						Short: true,
					},
				},
			},
			slack.SlackAttachment{
				Fallback: "View project: " + url,
				Color:    "good",
				Actions: []slack.SlackAction{
					slack.SlackAction{
						Type: "button",
						Text: "View project",
						URL:  url,
					},
				},
			},
			slack.SlackAttachment{
				Title:      "Kindly perform QA for this project.",
				Fallback:   "Kindly perform QA for this project.",
				CallbackID: "QA Response",
				Actions: []slack.SlackAction{
					slack.SlackAction{
						Type:  "button",
						Text:  "Approve",
						Name:  "approve",
						Value: string(marshaledPayload),
						Style: "primary",
					},
					slack.SlackAction{
						Type:  "button",
						Text:  "Reject",
						Name:  "reject",
						Value: string(marshaledPayload),
						Style: "danger",
					},
				},
			},
		},
	}

	return message, nil
}

func getQaSlackAttachment(build Build) slack.SlackAttachment {

	qaTeamAttachment := slack.SlackAttachment{
		Title:    "QA to be done by:",
		Fallback: "QA to be done by:",
	}

	for _, user := range build.Project.QA {
		qaTeamAttachment.Fields = append(qaTeamAttachment.Fields, slack.SlackField{
			Value: "<@" + user + ">",
			Short: true,
		})
	}

	return qaTeamAttachment
}

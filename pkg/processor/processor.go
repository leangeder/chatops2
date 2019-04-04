package processor

import (
	"encoding/json"
	"log"

	"github.com/leangeder/chatops2/pkg/server"
	"github.com/leangeder/chatops2/pkg/slack"
)

type actionPayload struct {
	// Build         Build      `json:"build,omitempty"`
	OwnerMessages []ownerMsg `json:"owner_messages,omitempty"`
}

type ownerMsg struct {
	Owner   string `json:"owner,omitempty"`
	Ts      string `json:"ts,omitempty"`
	Channel string `json:"channel,omitempty"`
}

// func (s *server.Server) startProcessors() {
// 	go s.buildProcessor()
// 	go s.interactionProcessor()
func startProcessors() {
	go buildProcessor()
	go interactionProcessor()
}

func (s *server.Server) buildProcessor() {
	for build := range s.Builds {
		go func() {

			ts, attemptErr := AttemptDeployMessage(build)
			if attemptErr != nil {
				log.Println(attemptErr)
				return
			}

			url, deployErr := deploy(build)

			if deployErr != nil {
				log.Println(deployErr)

				failErr := FailedDeployMessage(build, ts, deployErr)
				if failErr != nil {
					log.Println(failErr)
				}
				return
			}

			err := DeploySuccessMessage(build, ts, url)
			if err != nil {
				log.Println(err)
				return
			}

			payload, errs := OwnerMessages(build, url)
			if len(errs) > 0 {
				log.Println(errs)
				return
			}

			errs = QaMessages(build, url, payload)
			if len(errs) > 0 {
				log.Println(errs)
				return
			}

		}()
	}
}

func (s *server.Server) interactionProcessor() {
	for interaction := range s.Interactions {
		go func() {
			switch interaction.CallbackID {
			case "QA Response":
				go handleQaResponse(interaction)
			case "Deploy Decision":
				go handleOwnerDeploy(interaction)
			}
		}()
	}
}

func handleQaResponse(action slack.SlackInteraction) {

	user := action.User["id"]
	channel := action.Channel["id"]

	var payload actionPayload
	err := json.Unmarshal([]byte(action.Actions[0].Value), &payload)
	if err != nil {
		log.Println(err)
		return
	}

	var newAttch slack.SlackAttachment

	switch action.Actions[0].Name {
	case "approve":
		newAttch = slack.SlackAttachment{
			Title:    "Approved",
			Fallback: "Approved",
			Color:    "good",
		}
	case "reject":
		newAttch = slack.SlackAttachment{
			Title:    "Rejected",
			Fallback: "Rejected",
			Color:    "danger",
		}
	}

	updtMsg := action.OrigMessage
	updtMsg.Channel = channel
	updtMsg.Ts = action.MessageTs
	updtMsg.Update = true
	updtMsg.Attachments = updtMsg.Attachments[:2]
	updtMsg.Attachments = append(updtMsg.Attachments, newAttch)

	_, err = slack.sendSlack(updtMsg)
	if err != nil {
		log.Println(err)
		return
	}

	// Create a new slack message and add it as a threaded
	// reply to the Owner messages
	var newM slack.SlackMessage
	newM.Text = "<@" + user + "> has *" + newAttch.Title + "* this build"

	for _, oM := range payload.OwnerMessages {
		newM.ThreadTs = oM.Ts
		newM.Channel = oM.Channel

		_, err = slack.sendSlack(newM)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	return
}

func handleOwnerDeploy(action slack.SlackInteraction) {

	var errs []error

	switch action.Actions[0].Name {
	case "deploy":
		errs = slack.handleDeployToProd(action)
	case "close":
		errs = slack.handleCloseDeployment(action)
	}

	if len(errs) > 0 {
		log.Println(errs)
		return
	}

	return
}

func handleDeployToProd(action slack.SlackInteraction) (errs []error) {

	var payload actionPayload
	err := json.Unmarshal([]byte(action.Actions[0].Value), &payload)
	if err != nil {
		errs = append(errs, err)
		return
	}

	url, deployErr := slack.deployToProd(payload.Build)

	if deployErr != nil {
		errs = append(errs, deployErr)

		failErr := slack.sendFailedProdDeploy(payload, deployErr)
		if failErr != nil {
			errs = append(errs, failErr)
		}
		return
	}

	updtMsg := action.OrigMessage
	updtMsg.Update = true
	updtMsg.Attachments = updtMsg.Attachments[:3]
	updtMsg.Attachments = append(updtMsg.Attachments,
		slack.SlackAttachment{
			Title:    "Deployed to production",
			Fallback: "Deployed to production",
			Color:    "good",
		},
	)

	for _, oM := range payload.OwnerMessages {
		updtMsg.Channel = oM.Channel
		updtMsg.Ts = oM.Ts

		_, err = slack.sendSlack(updtMsg)
		if err != nil {
			errs = append(errs, err)
			continue
		}
	}

	err = slack.sendSuccessProdDeploy(payload, action.User["id"], url)
	if err != nil {
		errs = append(errs, err)
	}

	return
}

func handleCloseDeployment(action slack.SlackInteraction) (errs []error) {

	var payload actionPayload
	err := json.Unmarshal([]byte(action.Actions[0].Value), &payload)
	if err != nil {
		errs = append(errs, err)
		return
	}

	updateMessage := action.OrigMessage
	updateMessage.Update = true
	updateMessage.Attachments = updateMessage.Attachments[:3]
	updateMessage.Attachments = append(updateMessage.Attachments,
		slack.SlackAttachment{
			Title:    "Closed",
			Fallback: "Closed",
			Color:    "danger",
		},
	)

	for _, oM := range payload.OwnerMessages {
		updateMessage.Channel = oM.Channel
		updateMessage.Ts = oM.Ts

		_, err = slack.sendSlack(updateMessage)
		if err != nil {
			errs = append(errs, err)
			continue
		}
	}

	return
}

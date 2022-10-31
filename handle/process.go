package handle

import (
	"fmt"
	"log"

	"codeforces/boot"
	"codeforces/utils"
)

type Process struct {
	Broadcasts  []Broadcast
	MessageChan chan utils.GroupMessage
}

func (p *Process) InitBroadcast(broadcast Broadcast) {
	submissions, err := broadcast.QuerySubmission()
	if err != nil || len(submissions) == 0 {
		fmt.Println(err, len(submissions))
		panic(broadcast.Name() + " InitBroadcast err")
	}
	broadcast.ResetLastSubmissionID(submissions[0].ID)
}

func (p *Process) Refresh(broadcast Broadcast) []utils.GroupMessage {
	submissions, err := broadcast.QuerySubmission()
	if err != nil {
		log.Printf("%s Refresh QuerySubmission error: %v", broadcast.Name(), err)
		return nil
	}
	resp := make([]utils.GroupMessage, 0)
	for i := len(submissions) - 1; i >= 0; i-- {
		broadcast.ResetLastSubmissionID(submissions[i].ID)
		if submissions[i].Verdict == "OK" {
			resp = append(resp, utils.GroupMessage{
				GroupID: broadcast.GetGroupID(),
				Message: broadcast.AcceptedMessage(submissions[i]),
			})
		} else {
			resp = append(resp, utils.GroupMessage{
				GroupID: broadcast.GetGroupID(),
				Message: broadcast.WrongAnswerMessage(submissions[i]),
			})
		}
	}
	return resp
}

func (p *Process) ProcessMessage() {
	for {
		message := <-p.MessageChan

		err := boot.SendGroupMsg(message.GroupID, message.Message)
		log.Println(err, message.Message, message.GroupID)
	}
}

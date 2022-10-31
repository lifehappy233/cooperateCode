package handle

import "codeforces/utils"

type Broadcast interface {
	QuerySubmission() ([]utils.SubmissionInfo, error)
	WrongAnswerMessage(utils.SubmissionInfo) string
	AcceptedMessage(utils.SubmissionInfo) string
	ResetLastSubmissionID(int64)
	GetGroupID() string
	Name() string
}

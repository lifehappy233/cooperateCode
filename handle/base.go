package handle

import (
	"codeforces/crawl"
	"codeforces/utils"
)

type Base struct {
	crawl.Crawl
	LastSubmissionID int64
	GroupID          string
}

func (b *Base) QuerySubmission() ([]utils.SubmissionInfo, error) {
	return b.Crawl.QuerySubmission(b.LastSubmissionID)
}

func (b *Base) ResetLastSubmissionID(submissionID int64) {
	b.LastSubmissionID = submissionID
}

func (b *Base) GetGroupID() string {
	return b.GroupID
}

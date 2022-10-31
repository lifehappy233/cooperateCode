package handle

import (
	"fmt"

	"codeforces/crawl"
	"codeforces/utils"
)

type DefaultBroadcast struct {
	Base
}

func (d *DefaultBroadcast) WrongAnswerMessage(info utils.SubmissionInfo) string {
	return fmt.Sprintf("%s 刚刚 WA %s  %s了", d.Crawl.CodeforcesName, info.ContestName, info.ProblemName)
}

func (d *DefaultBroadcast) AcceptedMessage(info utils.SubmissionInfo) string {
	return fmt.Sprintf("%s 刚刚 AC 了 %s  %s", d.Crawl.CodeforcesName, info.ContestName, info.ProblemName)
}

func (d *DefaultBroadcast) Name() string {
	return "DefaultBroadcast"
}

func NewDefaultBroadcast(codeforcesName, groupID string) DefaultBroadcast {
	return DefaultBroadcast{
		Base: Base{
			GroupID: groupID,
			Crawl: crawl.Crawl{
				CodeforcesName: codeforcesName,
			},
		},
	}
}

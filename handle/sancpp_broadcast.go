package handle

import (
	"codeforces/crawl"
	"fmt"

	"codeforces/utils"
)

type SancppBroadcast struct {
	Base
}

func (s *SancppBroadcast) WrongAnswerMessage(info utils.SubmissionInfo) string {
	return fmt.Sprintf("小甜甜 被 %s  %s 击倒了[CQ:face,id=35][CQ:face,id=35][CQ:face,id=35]", info.ContestName, info.ProblemName)
}

func (s *SancppBroadcast) AcceptedMessage(info utils.SubmissionInfo) string {
	return fmt.Sprintf("小甜甜 击杀了 %s  %s[CQ:face,id=2][CQ:face,id=2][CQ:face,id=2]", info.ContestName, info.ProblemName)
}

func (s *SancppBroadcast) Name() string {
	return "SancppBroadcast"
}

func NewSancppBroadcast(groupID string) SancppBroadcast {
	return SancppBroadcast{
		Base: Base{
			GroupID: groupID,
			Crawl: crawl.Crawl{
				CodeforcesName: "sancpp",
			},
		},
	}
}

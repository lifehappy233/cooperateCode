package handle

import (
	"codeforces/crawl"
	"fmt"

	"codeforces/utils"
)

type LifehappyBroadcast struct {
	Base
}

func (s *LifehappyBroadcast) WrongAnswerMessage(info utils.SubmissionInfo) string {
	return fmt.Sprintf("康神 误入假题 %s  %s [CQ:face,id=35][CQ:face,id=35][CQ:face,id=35]", info.ContestName, info.ProblemName)
}

func (s *LifehappyBroadcast) AcceptedMessage(info utils.SubmissionInfo) string {
	return fmt.Sprintf("康神 击杀了 %s  %s[CQ:face,id=2][CQ:face,id=2][CQ:face,id=2]", info.ContestName, info.ProblemName)
}

func (s *LifehappyBroadcast) Name() string {
	return "LifehappyBroadcast"
}

func NewLifehappyBroadcast(groupID string) LifehappyBroadcast {
	return LifehappyBroadcast{
		Base: Base{
			GroupID: groupID,
			Crawl: crawl.Crawl{
				CodeforcesName: "lifehappy01",
			},
		},
	}
}

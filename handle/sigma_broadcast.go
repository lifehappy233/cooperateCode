package handle

import (
	"codeforces/crawl"
	"fmt"

	"codeforces/utils"
)

type SigmaBroadcast struct {
	Base
}

func (s *SigmaBroadcast) WrongAnswerMessage(info utils.SubmissionInfo) string {
	return fmt.Sprintf("sigma 被 %s  %s 给爆杀啦[CQ:face,id=9][CQ:face,id=9][CQ:face,id=9]", info.ContestName, info.ProblemName)
}

func (s *SigmaBroadcast) AcceptedMessage(info utils.SubmissionInfo) string {
	return fmt.Sprintf("sigma大爹 刚刚乱杀了 %s  %s[CQ:face,id=144][CQ:face,id=144][CQ:face,id=144]", info.ContestName, info.ProblemName)
}

func (s *SigmaBroadcast) Name() string {
	return "SigmaBroadcast"
}

func NewSigmaBroadcast(groupID string) SigmaBroadcast {
	return SigmaBroadcast{
		Base: Base{
			GroupID: groupID,
			Crawl: crawl.Crawl{
				CodeforcesName: "laysan",
			},
		},
	}
}

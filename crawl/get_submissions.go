package crawl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

// QuerySubmission return (verdict, contest name, problem name)
func QuerySubmission(submissionID int64) ([]SubmissionInfo, error) {
	res, err := client.Do(submissionReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	rsp := &ParseSubmission{}
	err = json.Unmarshal(body, rsp)
	if err != nil {
		return nil, err
	}
	if len(rsp.Result) == 0 {
		return nil, fmt.Errorf("query result list empty")
	}

	resp := make([]SubmissionInfo, 0)
	for _, result := range rsp.Result {
		if result.Verdict == "TESTING" {
			continue
		}
		if result.ID != submissionID {
			submissionInfo := SubmissionInfo{
				ID:          result.ID,
				Verdict:     result.Verdict,
				ProblemName: result.Problem.Name,
			}
			contestID := result.Problem.ContestID

			if contestName, have := contestID2Name[contestID]; have {
				submissionInfo.ContestName = contestName
			} else {
				contestName, err = QueryContest(strconv.FormatInt(contestID, 10))
				if err != nil {
					return nil, err
				}
				contestID2Name[contestID] = contestName
				submissionInfo.ContestName = contestName
			}
			resp = append(resp, submissionInfo)
		} else {
			break
		}
	}
	return resp, nil
}

type SubmissionInfo struct {
	ID          int64
	Verdict     string
	ProblemName string
	ContestName string
}

type ParseSubmission struct {
	Status string       `json:"status"`
	Result []Submission `json:"result"`
}

type Submission struct {
	ID      int64   `json:"id"`
	Verdict string  `json:"verdict"`
	Problem Problem `jsoN:"problem"`
}

type Problem struct {
	ContestID int64  `json:"contestId"`
	Index     string `json:"index"`
	Name      string `json:"name"`
}

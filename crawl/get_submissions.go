package crawl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"codeforces/utils"
)

// QuerySubmission return (verdict, contest name, problem name)
func (c *Crawl) QuerySubmission(submissionID int64) ([]utils.SubmissionInfo, error) {
	c.submissionReqOnce.Do(c.initSubmissionReq)

	res, err := client.Do(c.submissionReq)
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

	resp := make([]utils.SubmissionInfo, 0)
	for _, result := range rsp.Result {
		if result.Verdict == "TESTING" {
			continue
		}
		if result.ID != submissionID {
			submissionInfo := utils.SubmissionInfo{
				ID:          result.ID,
				Verdict:     result.Verdict,
				ProblemName: result.Problem.Index + ". " + result.Problem.Name,
			}
			contestID := result.Problem.ContestID

			if contestName, have := contestID2Name[contestID]; have {
				submissionInfo.ContestName = contestName
			} else {
				contestName, err = c.QueryContest(strconv.FormatInt(contestID, 10))
				if err != nil {
					return nil, err
				}
				log.Println("new context record:", contestID, contestName)
				contestID2Name[contestID] = contestName
				submissionInfo.ContestName = contestName
			}
			resp = append(resp, submissionInfo)
			log.Println(submissionInfo.ContestName + "  " + submissionInfo.ProblemName)
		} else {
			break
		}
	}
	return resp, nil
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

func (c *Crawl) initSubmissionReq() {
	url := "https://codeforces.com/api/user.status?handle=" + c.CodeforcesName + "&from=1&count=10"
	method := "GET"
	var err error
	c.submissionReq, err = http.NewRequest(method, url, nil)

	if err != nil {
		panic(err)
	}
	c.submissionReq.Header.Add("authority", "codeforces.com")
	c.submissionReq.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	c.submissionReq.Header.Add("accept-language", "zh-CN,zh;q=0.9,en-CN;q=0.8,en;q=0.7")
	c.submissionReq.Header.Add("cache-control", "max-age=0")
	c.submissionReq.Header.Add("cookie", "__utmc=71512449; evercookie_png=u11lhwgv8qp1275nbs; evercookie_etag=u11lhwgv8qp1275nbs; evercookie_cache=u11lhwgv8qp1275nbs; 70a7c28f3de=u11lhwgv8qp1275nbs; JSESSIONID=E8671ED7EB4D2856F68AA16A217669C4-n1; 39ce7=CFhY24ou; _ga=GA1.2.365646997.1665840576; X-User=; lastOnlineTimeUpdaterInvocation=1667018126805; X-User-Sha1=eefd78d3e2b842abcc5cd0a8e63c1d694dafeed1; __utma=71512449.365646997.1665840576.1667016035.1667020231.21; __utmz=71512449.1667020231.21.2.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __utmb=71512449.28.10.1667020231")
	c.submissionReq.Header.Add("sec-ch-ua", "\"Chromium\";v=\"106\", \"Google Chrome\";v=\"106\", \"Not;A=Brand\";v=\"99\"")
	c.submissionReq.Header.Add("sec-ch-ua-mobile", "?0")
	c.submissionReq.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	c.submissionReq.Header.Add("sec-fetch-dest", "document")
	c.submissionReq.Header.Add("sec-fetch-mode", "navigate")
	c.submissionReq.Header.Add("sec-fetch-site", "none")
	c.submissionReq.Header.Add("sec-fetch-user", "?1")
	c.submissionReq.Header.Add("upgrade-insecure-requests", "1")
	c.submissionReq.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
}

package crawl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func QueryContest(contestID string) (string, error) {
	err := initContestReq(contestID)
	if err != nil {
		return "", err
	}

	res, err := client.Do(contestReq)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	parseContest := &ParseContest{}
	err = json.Unmarshal(body, parseContest)
	if err != nil {
		return "", err
	}
	if parseContest.Result.Contest.Name == "" {
		return "", fmt.Errorf("have not find contest name")
	}

	return parseContest.Result.Contest.Name, nil
}

type ParseContest struct {
	Result ContestResult `json:"result"`
}

type ContestResult struct {
	Contest Contest `json:"contest"`
}

type Contest struct {
	Name string `json:"name"`
}

package crawl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Crawl) QueryContest(contestID string) (string, error) {
	err := c.initContestReq(contestID)
	if err != nil {
		return "", err
	}

	res, err := client.Do(c.contestReq)
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

func (c *Crawl) initContestReq(contestID string) error {
	url := "https://codeforces.com/api/contest.standings?contestId=" + contestID + "&from=1&count=1&showUnofficial=true"
	method := "GET"

	var err error
	c.contestReq, err = http.NewRequest(method, url, nil)

	if err != nil {
		return err
	}
	c.contestReq.Header.Add("authority", "codeforces.com")
	c.contestReq.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	c.contestReq.Header.Add("accept-language", "zh-CN,zh;q=0.9,en-CN;q=0.8,en;q=0.7")
	c.contestReq.Header.Add("cache-control", "max-age=0")
	c.contestReq.Header.Add("cookie", "__utmc=71512449; evercookie_png=u11lhwgv8qp1275nbs; evercookie_etag=u11lhwgv8qp1275nbs; evercookie_cache=u11lhwgv8qp1275nbs; 70a7c28f3de=u11lhwgv8qp1275nbs; JSESSIONID=E8671ED7EB4D2856F68AA16A217669C4-n1; 39ce7=CFhY24ou; _ga=GA1.2.365646997.1665840576; X-User=; lastOnlineTimeUpdaterInvocation=1667018126805; X-User-Sha1=eefd78d3e2b842abcc5cd0a8e63c1d694dafeed1; __utmz=71512449.1667020231.21.2.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __utma=71512449.365646997.1665840576.1667020231.1667026476.22; __utmt=1; __utmb=71512449.5.10.1667026476")
	c.contestReq.Header.Add("sec-ch-ua", "\"Chromium\";v=\"106\", \"Google Chrome\";v=\"106\", \"Not;A=Brand\";v=\"99\"")
	c.contestReq.Header.Add("sec-ch-ua-mobile", "?0")
	c.contestReq.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	c.contestReq.Header.Add("sec-fetch-dest", "document")
	c.contestReq.Header.Add("sec-fetch-mode", "navigate")
	c.contestReq.Header.Add("sec-fetch-site", "none")
	c.contestReq.Header.Add("sec-fetch-user", "?1")
	c.contestReq.Header.Add("upgrade-insecure-requests", "1")
	c.contestReq.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")

	return nil
}

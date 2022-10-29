package crawl

import (
	"net/http"
)

var (
	client         *http.Client
	submissionReq  *http.Request
	contestReq     *http.Request
	contestID2Name map[int64]string
)

func init() {
	client = &http.Client{}
	initSubmissionReq()
	contestID2Name = make(map[int64]string)
}

func initSubmissionReq() {
	url := "https://codeforces.com/api/user.status?handle=laysan&from=1&count=5"
	method := "GET"
	var err error
	submissionReq, err = http.NewRequest(method, url, nil)

	if err != nil {
		panic(err)
	}
	submissionReq.Header.Add("authority", "codeforces.com")
	submissionReq.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	submissionReq.Header.Add("accept-language", "zh-CN,zh;q=0.9,en-CN;q=0.8,en;q=0.7")
	submissionReq.Header.Add("cache-control", "max-age=0")
	submissionReq.Header.Add("cookie", "__utmc=71512449; evercookie_png=u11lhwgv8qp1275nbs; evercookie_etag=u11lhwgv8qp1275nbs; evercookie_cache=u11lhwgv8qp1275nbs; 70a7c28f3de=u11lhwgv8qp1275nbs; JSESSIONID=E8671ED7EB4D2856F68AA16A217669C4-n1; 39ce7=CFhY24ou; _ga=GA1.2.365646997.1665840576; X-User=; lastOnlineTimeUpdaterInvocation=1667018126805; X-User-Sha1=eefd78d3e2b842abcc5cd0a8e63c1d694dafeed1; __utma=71512449.365646997.1665840576.1667016035.1667020231.21; __utmz=71512449.1667020231.21.2.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __utmb=71512449.28.10.1667020231")
	submissionReq.Header.Add("sec-ch-ua", "\"Chromium\";v=\"106\", \"Google Chrome\";v=\"106\", \"Not;A=Brand\";v=\"99\"")
	submissionReq.Header.Add("sec-ch-ua-mobile", "?0")
	submissionReq.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	submissionReq.Header.Add("sec-fetch-dest", "document")
	submissionReq.Header.Add("sec-fetch-mode", "navigate")
	submissionReq.Header.Add("sec-fetch-site", "none")
	submissionReq.Header.Add("sec-fetch-user", "?1")
	submissionReq.Header.Add("upgrade-insecure-requests", "1")
	submissionReq.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
}

func initContestReq(contestID string) error {
	url := "https://codeforces.com/api/contest.standings?contestId=" + contestID + "&from=1&count=1&showUnofficial=true"
	method := "GET"

	var err error
	contestReq, err = http.NewRequest(method, url, nil)

	if err != nil {
		return err
	}
	contestReq.Header.Add("authority", "codeforces.com")
	contestReq.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	contestReq.Header.Add("accept-language", "zh-CN,zh;q=0.9,en-CN;q=0.8,en;q=0.7")
	contestReq.Header.Add("cache-control", "max-age=0")
	contestReq.Header.Add("cookie", "__utmc=71512449; evercookie_png=u11lhwgv8qp1275nbs; evercookie_etag=u11lhwgv8qp1275nbs; evercookie_cache=u11lhwgv8qp1275nbs; 70a7c28f3de=u11lhwgv8qp1275nbs; JSESSIONID=E8671ED7EB4D2856F68AA16A217669C4-n1; 39ce7=CFhY24ou; _ga=GA1.2.365646997.1665840576; X-User=; lastOnlineTimeUpdaterInvocation=1667018126805; X-User-Sha1=eefd78d3e2b842abcc5cd0a8e63c1d694dafeed1; __utmz=71512449.1667020231.21.2.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __utma=71512449.365646997.1665840576.1667020231.1667026476.22; __utmt=1; __utmb=71512449.5.10.1667026476")
	contestReq.Header.Add("sec-ch-ua", "\"Chromium\";v=\"106\", \"Google Chrome\";v=\"106\", \"Not;A=Brand\";v=\"99\"")
	contestReq.Header.Add("sec-ch-ua-mobile", "?0")
	contestReq.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	contestReq.Header.Add("sec-fetch-dest", "document")
	contestReq.Header.Add("sec-fetch-mode", "navigate")
	contestReq.Header.Add("sec-fetch-site", "none")
	contestReq.Header.Add("sec-fetch-user", "?1")
	contestReq.Header.Add("upgrade-insecure-requests", "1")
	contestReq.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")

	return nil
}

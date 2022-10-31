package crawl

import (
	"net/http"
	"sync"
)

var (
	client         *http.Client
	contestID2Name map[int64]string
)

func init() {
	client = &http.Client{}
	contestID2Name = make(map[int64]string)
}

type Crawl struct {
	CodeforcesName    string
	submissionReq     *http.Request
	submissionReqOnce sync.Once
	contestReq        *http.Request
}

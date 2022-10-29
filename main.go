package main

import (
	"codeforces/boot"
	"codeforces/crawl"
	"fmt"
	"time"
)

func handle() {
	submissions, err := crawl.QuerySubmission(0)
	//fmt.Println(crawl.QueryContest("102992"))
	if err != nil || len(submissions) == 0 {
		fmt.Println(err, len(submissions))
		panic("err")
	}
	preSubmissionID := submissions[0].ID
	fmt.Println(preSubmissionID)

	submissionChan := make(chan crawl.SubmissionInfo, 10)

	go func() {
		for {
			submission := <-submissionChan
			errSend := boot.SendGroupMsg("770539963", submission.Verdict+submission.ContestName+submission.ProblemName)
			fmt.Println(errSend, submission.Verdict+submission.ContestName+submission.ProblemName)
		}
	}()

	for range time.Tick(5 * time.Second) {
		submissions, err = crawl.QuerySubmission(preSubmissionID)
		for i := len(submissions) - 1; i >= 0; i-- {
			submissionChan <- submissions[i]
			preSubmissionID = submissions[i].ID
		}
	}
}

func main() {
	//handle
	handle()
}

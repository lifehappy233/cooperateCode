package utils

type SubmissionInfo struct {
	ID          int64
	Verdict     string
	ProblemName string
	ContestName string
}

type GroupMessage struct {
	GroupID string
	Message string
}

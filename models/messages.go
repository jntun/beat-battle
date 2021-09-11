package models

import "time"

type VoteMessage struct {
	Timestamp time.Time `json:"-"`
	//User       `json:"user"`
	User       UserMsg `json:"user"`
	Submission string  `json:"submission"`
}

type UserMsg struct {
	Id   string
	Name string
}

type SubmissionMessage struct {
	Id       string
	Author   UserMsg `json:"author,omitempty"`
	Resource string
	Type     uint
}

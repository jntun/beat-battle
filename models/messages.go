package models

import "time"

type VoteMessage struct {
	Timestamp  time.Time `json:"-"`
	User       UserMsg   `json:"user"`
	Submission string    `json:"submission"`
}

type UserMsg struct {
	Id   string
	Name string
}

type SubmissionMessage struct {
	Id       string `json:",omitempty"`
	Author   UserMsg
	Resource string
	Type     uint
}

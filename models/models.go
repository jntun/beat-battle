package models

import (
	"net/url"
	"sync"
	"time"

	"github.com/satori/go.uuid"
)

type ID struct {
	uuid.UUID
}

type Battle struct {
	ID
	Title            string
	Sequence         uint
	Submissions      map[string]*Submission
	SubLock          sync.Mutex
	submissionWindow TimeWindow
	voteWindow       TimeWindow
}

type Submission struct {
	ID       `json:"id,omitempty"`
	Author   User    `json:"author"`
	Resource url.URL `json:"resource"`
	Votes    uint    `json:"votes,omitempty"`
	Type     uint    `json:"type"`
}

type User struct {
	ID
	Name string `json:"username"`
}

type TimeWindow struct {
	start time.Time
	end   time.Time
}

func NewBattle() Battle {
	subMap := make(map[string]*Submission)
	return Battle{
		newID(),
		"Test battle #1",
		0,
		subMap,
		sync.Mutex{},
		TimeWindow{
			start: time.Now(),
			end:   time.Now().Add(time.Hour),
		},
		TimeWindow{
			start: time.Now().Add(time.Hour * 2),
			end:   time.Now().Add(time.Hour * 3),
		},
	}
}

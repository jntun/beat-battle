package models

import (
	"fmt"
)

func (sub Submission) String() string {
	return fmt.Sprintf("subm %s: %s %s - %d", sub.ID.String()[:8], sub.Author, sub.Resource.String(), sub.Type)
}

func (user User) String() string {
	return fmt.Sprintf("%s#%s", user.UUID.String()[:8], user.Name)
}

func (sub SubmissionMessage) String() string {
	return fmt.Sprintf("%s: %s - %d", sub.Author, sub.Resource, sub.Type)
}

func (vote VoteMessage) String() string {
	if len(vote.Submission) < 8 {
		return fmt.Sprintf("vote %s: %s", vote.User, vote.Submission)
	}
	return fmt.Sprintf("vote %s: %s", vote.User, vote.Submission[:8])
}

func (user UserMsg) String() string {
	return fmt.Sprintf("%s#%s", user.Id[:8], user.Name)
}

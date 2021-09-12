package models

import (
	"fmt"
)

func (sub Submission) String() string {
	return fmt.Sprintf("subm %s: %s %s - %d", sub.Author, sub.ID.String(), sub.Resource.String(), sub.Type)
}

func (user User) String() string {
	return fmt.Sprintf("%s '%s'", user.UUID.String()[:8], user.Name)
}

func (sub SubmissionMessage) String() string {
	return fmt.Sprintf("%s: %s - %d", sub.Author, sub.Resource, sub.Type)
}

func (vote VoteMessage) String() string {
	return fmt.Sprintf("vote %s: %s", vote.User, vote.Submission[:8])
}

func (user UserMsg) String() string {
	return fmt.Sprintf("%s#%s", user.Id[:8], user.Name)
}

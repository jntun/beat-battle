package models

import (
	uuid "github.com/satori/go.uuid"
	"net/url"
)

func (msg SubmissionMessage) ToSubmission() (*Submission, error) {
	resource, err := url.Parse(msg.Resource)
	if err != nil {
		return nil, err
	}

	var user *User
	if user, err = msg.Author.ToUser(); err != nil {
		return nil, err
	}

	sub := Submission{
		newID(),
		*user,
		*resource,
		0,
		msg.Type,
	}

	return &sub, nil
}

func (msg UserMsg) ToUser() (*User, error) {
	var (
		err error
		id  *ID
	)
	if id, err = IDfrom(msg.Id); err != nil {
		return nil, err
	}

	return &User{
		ID:   *id,
		Name: msg.Name,
	}, nil
}

func newID() ID {
	return ID{uuid.NewV4()}
}

func IDfrom(s string) (*ID, error) {
	uid, err := uuid.FromString(s)
	if err != nil {
		return nil, err
	}
	return &ID{uid}, nil
}

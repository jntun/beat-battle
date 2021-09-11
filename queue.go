package main

import (
	"beat-battle/models"
	"fmt"
	"log"
)

func (sess *Session) drainSubmitQueue() error {
	for _, subMsg := range sess.submissionQueue {
		submission, err := subMsg.ToSubmission()
		if err != nil {
			/* TODO process and see if we want to continue */
			return fmt.Errorf("%v, %s", subMsg, err)
		}
		log.Printf("draining: %v | ref: %p\n", *submission, submission)
		sess.battle.Submissions[submission.UUID.String()] = *submission
	}
	return nil
}

func (sess *Session) drainVoteQueue() error {
	for _, voteMsg := range sess.voteQueue {
		if verifyVote(voteMsg) {
			if err := sess.processVote(voteMsg); err != nil {
				return err
			}
		}
	}
	return nil
}

func verifyVote(vote models.VoteMessage) bool {
	/* TODO logically verify vote message */
	return true
}

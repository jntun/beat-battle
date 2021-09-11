package main

import (
	"beat-battle/models"
	"fmt"
	"net/http"
	"strings"
)

const port = 8000
const endpoint = "/battle/"

type Session struct {
	battle          models.Battle
	submissionQueue []models.SubmissionMessage
	voteQueue       []models.VoteMessage
}

func (sess *Session) hookHttpServer() {
	http.HandleFunc(endpoint, sess.main)
}

func (sess Session) getAddr() string {
	return fmt.Sprintf(":%d", port)
}

func (sess *Session) main(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	genLog("Main", r.URL.String())
	splitURL := strings.Split(r.URL.String(), "/")
	if !(len(splitURL) > 1) {
		// for only /battle/ endpoint
		return
	}
	switch splitURL[2] {
	case "vote":
		sess.AddVote(w, r)
		sess.tryVoteDrain()
	case "submissions":
		sess.GetSubmissions(w, r)
	case "submit":
		sess.Submit(w, r)
		sess.tryDrain()
	}

}

func (sess *Session) tryDrain() {
	if err := sess.drainSubmitQueue(); err != nil {
		subLog("couldn't drain sub queue: %s", err)
	}
	sess.tryVoteDrain()
}

func (sess *Session) tryVoteDrain() {
	if err := sess.drainVoteQueue(); err != nil {
		subLog("couldn't drain vote queue: %s", err)
	}
}

func (sess *Session) processVote(vote models.VoteMessage) error {
	submission, found := sess.battle.Submissions[vote.Submission]
	if !found {
		return fmt.Errorf("could not find submission matching uuid '%s'", vote.Submission)
	}
	submission.Votes++
	return nil
}

func NewSession() Session {
	return Session{models.NewBattle(), make([]models.SubmissionMessage, 0), make([]models.VoteMessage, 0)}
}

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
	queueStat
}

func (sess *Session) hookHttpServer() {
	http.HandleFunc(endpoint, sess.main)
}

func (sess Session) getAddr() string {
	return fmt.Sprintf(":%d", port)
}

func (sess *Session) main(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	splitURL := strings.Split(r.URL.String(), "/")
	if !(len(splitURL) > 1) {
		// for only /battle/ endpoint
		return
	}
	switch splitURL[2] {
	case "vote":
		if r.Method == http.MethodPost {
			sess.AddVote(w, r)
			sess.tryVoteDrain()
		}
	case "submissions":
		sess.GetSubmissions(w, r)
	case "submit":
		if r.Method == http.MethodPut {
			sess.Submit(w, r)
			sess.tryDrain()
		}
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

func NewSession() Session {
	return Session{models.NewBattle(), submQueue(), voteQueue(), queueStat{queueIndex{0, 0, voteEntryPoint, voteHWM}, queueIndex{0, 0, submEntryPoint, submHWM}}}
}

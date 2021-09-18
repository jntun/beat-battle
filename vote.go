package main

import (
	"beat-battle/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (sess *Session) AddVote(w http.ResponseWriter, r *http.Request) {
	bodyMsg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeErrorResponse(w, err)
		return
	}

	if len(bodyMsg) == 0 {
		voteLog("body length was 0. rejecting message.")
		w.Write([]byte("Bad vote message."))
		return
	}

	voteMsg := models.VoteMessage{}
	err = json.Unmarshal(bodyMsg, &voteMsg)
	if err != nil {
		handleVoteError(w, err)
		return
	}

	voteMsg.Timestamp = time.Now()

	//voteLog("received: %v", voteMsg)
	sess.voteQueue[sess.queueStat.vote.length] = voteMsg
	sess.queueStat.vote.length++
	_, err = w.Write(models.ResponseStatus{
		Code:    0,
		Success: true,
	}.AsJSONBytes())
	if err != nil {
		voteError(err)
	}
}

func (sess *Session) processVote(vote models.VoteMessage) error {
	submission, found := sess.battle.Submissions[vote.Submission]
	if !found {
		return fmt.Errorf("could not find submission matching uuid '%s'", vote.Submission)
	}
	submission.Votes++
	//voteLog("adding vote for %v | %v | vc: %d", submission, vote, submission.Votes)
	return nil
}

func verifyVote(vote models.VoteMessage) bool {
	/* TODO logically verify vote message */
	return true
}

func handleVoteError(w http.ResponseWriter, err error) {
	voteError(err)
	writeErrorResponse(w, err)
}

func writeErrorResponse(w http.ResponseWriter, err error) {
	log.Println(err)
	rc, err := w.Write([]byte("Bad vote message."))
	if err != nil {
		log.Printf("[%d] Failure writing vote error response: %s.\n", rc, err)
	}
}

func voteError(err error) {
	voteLog("[ERROR] %v", err)
}

func voteLog(fmtStr string, args ...interface{}) {
	genLog("VOTE", fmtStr, args...)
}

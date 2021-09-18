package main

import (
	"beat-battle/models"
	"bytes"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSession_AddVoteLiveHWM(t *testing.T) {
	targetID, err := getSubmissionTarget()
	if err != nil {
		t.Error(err)
	}
	t.Log("target:", targetID)
	msg := models.VoteMessage{
		User: models.UserMsg{
			Id:   uuid.NewV4().String(),
			Name: "vote_user",
		},
		Submission: targetID,
	}
	jsonBody, _ := json.Marshal(msg)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8000/battle/vote", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Errorf("failed to forge vote request: %s", err)
	}

	client := http.Client{}
	// Send enough to hit the HWM point
	for i := 0; i < int(voteHWM); i++ {
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("failed vote request:")
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("failed vote body read: %s", err)
		}
		t.Log(string(body))
	}
}

func TestSession_AddVote(t *testing.T) {
	sess := NewSession()

	for i := 0; i < int(sess.queueStat.subm.entryPoint); i++ {
		req := makeSendReq(&sess)
		tryReq(t, &sess, req)
	}

	for i := 0; i < int(sess.queueStat.vote.entryPoint); i++ {
		req := makeVoteReq(&sess)
		tryReq(t, &sess, req)
	}
}

func makeVoteReq(sess *Session) *http.Request {
	var target *models.Submission
	for _, sub := range sess.battle.Submissions {
		target = sub
		break
	}
	msg := models.VoteMessage{
		User: models.UserMsg{
			Id:   "vote_user",
			Name: uuid.NewV4().String(),
		},
		Submission: target.ID.String(),
	}
	jsonBuf, _ := json.Marshal(msg)
	req := httptest.NewRequest(http.MethodPost, "http://localhost"+sess.getAddr()+endpoint+"vote", bytes.NewBuffer(jsonBuf))
	return req
}

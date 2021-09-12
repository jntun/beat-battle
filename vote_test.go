package main

import (
	"beat-battle/models"
	"bytes"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSession_AddVoteLive(t *testing.T) {
	t.Run("Live", TestSession_SubmitLive)

	tResp, err := http.Get("http://localhost:8000/battle/submissions")
	if err != nil {
		t.Errorf("couldn't get target body: %s", err)
	}
	tBody, err := ioutil.ReadAll(tResp.Body)
	if err != nil {
		t.Errorf("failed target response body read: %s", err)
	}
	var tObj interface{}
	err = json.Unmarshal(tBody, &tObj)
	var targetID string

	for _, key := range tObj.(map[string]interface{}) {
		targetID = key.(string)
		break
	}

	fmt.Println("xyxyxy:", targetID)

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

	for i := 0; i < int(sess.queueStat.submCount.entryPoint); i++ {
		req := makeSendReq(&sess)
		tryReq(t, &sess, req)
	}

	for i := 0; i < int(sess.queueStat.voteCount.entryPoint); i++ {
		req := makeVoteReq(&sess)
		tryReq(t, &sess, req)
	}
}

func makeVoteReq(sess *Session) *http.Request {
	var target models.Submission
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

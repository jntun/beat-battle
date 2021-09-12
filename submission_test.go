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

func TestSession_SubmitLive(t *testing.T) {
	msg := models.SubmissionMessage{
		Author: models.UserMsg{
			Id:   uuid.NewV4().String(),
			Name: "jntun",
		},
		Resource: "https://soundcloud.com/jntun/backflip",
		Type:     0,
	}

	jsonByte, _ := json.Marshal(msg)
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8000/battle/submit", bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Errorf("failed to forge subm request: %s", err)
	}

	client := http.Client{}
	for i := 0; i < int(submHWM); i++ {
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("failed subm request: %s", err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("failed response body read: %s", err)
		}
		t.Log(string(body))
	}

}

func TestSession_Submit(t *testing.T) {
	sess := NewSession()

	for i := 0; i < 10; i++ {
		req := makeSendReq(&sess)
		tryReq(t, &sess, req)
	}
}

func makeSendReq(sess *Session) *http.Request {
	msg := models.SubmissionMessage{
		Author: models.UserMsg{
			Id:   "362cfe38-4f4e-428d-aa2f-7bae392d9a99",
			Name: "jntun",
		},
		Resource: "https://soundcloud.com/jntun/backflip",
		Type:     0,
	}
	jsonBuf, _ := json.Marshal(msg)
	req := httptest.NewRequest(http.MethodPut, "http://localhost"+sess.getAddr()+endpoint+"submit", bytes.NewBuffer(jsonBuf))

	return req
}

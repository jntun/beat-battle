package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func tryReq(t *testing.T, sess *Session, r *http.Request) {
	w := httptest.NewRecorder()
	sess.main(w, r)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("couldn't read resp body: %s", err)
	}
	fmt.Printf("resp: %s\n", body)
}

func getSubmissionTarget() (string, error) {
	tResp, err := http.Get("http://localhost:8000/battle/submissions")
	if err != nil {
		return "", fmt.Errorf("couldn't get target body: %s", err)
	}
	tBody, err := ioutil.ReadAll(tResp.Body)
	if err != nil {
		return "", fmt.Errorf("failed target response body read: %s", err)
	}
	var tObj interface{}
	err = json.Unmarshal(tBody, &tObj)

	var targetID string
	for key, _ := range tObj.(map[string]interface{}) {
		targetID = key
		break
	}

	return targetID, nil
}
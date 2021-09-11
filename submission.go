package main

import (
	"beat-battle/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (sess *Session) Submit(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		subLog("%s", err)
		return
	}

	if len(body) == 0 {
		subLog("empty msg, skipping...")
		return
	}

	subMsg := &models.SubmissionMessage{}
	err = json.Unmarshal(body, subMsg)
	if err != nil {
		subLog("could not unpack submission msg: %s", err)
		return
	}

	subLog("new submission: %v", *subMsg)
	sess.submissionQueue = append(sess.submissionQueue, *subMsg)
}

func subLog(fmtStr string, args ...interface{}) {
	genLog("Subm", fmtStr, args)
}

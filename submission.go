package main

import (
	"beat-battle/models"
	"encoding/json"
	"io/ioutil"
	"log"
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

	subLog("received: %v", *subMsg)
	/* Insert into submission queue here in O(1) time since we pre-allocated all the slots */
	sess.submissionQueue[sess.queueStat.submCount.length] = *subMsg
	sess.queueStat.submCount.length++
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("{ \"status\": \"done\" } ")); err != nil {
		log.Printf("failed write: %s\n", err)
	}
}

func subLog(fmtStr string, args ...interface{}) {
	genLog("SUBM", fmtStr, args...)
}

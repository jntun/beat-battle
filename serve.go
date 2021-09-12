package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (sess *Session) GetSubmissions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Write([]byte(fmt.Sprintf("Bad method '%s'.", r.Method)))
		return
	}

	if len(sess.battle.Submissions) == 0 {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("No submissions."))
		return
	}

	sess.battle.SubLock.Lock()
	msg, err := json.Marshal(sess.battle.Submissions)
	if err != nil {
		serveLog("%v", err)
		return
	}
	sess.battle.SubLock.Unlock()

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(msg)
	if err != nil {
		serveLog("%v", err)
		return
	}
}

func serveLog(fmtStr string, args ...interface{}) {
	genLog("SERVE", fmtStr, args...)
}

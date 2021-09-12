package main

import (
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

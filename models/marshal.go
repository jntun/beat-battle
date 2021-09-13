package models

import (
	"bytes"
	"encoding/json"
	"io"
)

func (subm Submission) MarshalJSON() ([]byte, error) {
	buff := bytes.NewBuffer([]byte{})
	buff.Write([]byte("{"))
	//addKeyValTrail(buff, "uuid", subm.Author.ID.String())
	addKeyValTrail(buff, "votes", subm.Votes)
	addKeyValTrail(buff, "type", subm.Type)
	addKeyVal(buff, "resource", subm.Resource.String())
	buff.Write([]byte("}"))

	return buff.Bytes(), nil
}

func addKeyValTrail(w io.Writer, key string, val interface{}) {
	addKeyVal(w, key, val)
	w.Write([]byte(","))
}

func addKeyVal(w io.Writer, key string, val interface{}) {
	valBytes, err := json.Marshal(val)
	if err != nil {
		// fail
	}
	w.Write([]byte("\""))
	w.Write([]byte(key))
	w.Write([]byte("\":"))
	w.Write(valBytes)
}

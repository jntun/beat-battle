package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	sess := NewSession()
	log.Printf("Starting on %s...\n", sess.getAddr())
	server := &http.Server{Addr: sess.getAddr()}
	sess.hookHttpServer()
	log.Fatalln(server.ListenAndServe())
}

func genLog(prefix string, fmtStr string, args ...interface{}) {
	builder := strings.Builder{}
	builder.Write([]byte("["))
	builder.Write([]byte(prefix))
	builder.Write([]byte("] "))
	builder.Write([]byte(fmtStr))
	builder.Write([]byte("\n"))
	if len(args) != 0 {
		log.Printf(builder.String(), args)
	} else {
		log.Print(builder.String())
	}
}

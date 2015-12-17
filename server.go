package main

import (
	"log"
	"net/http"
	"os"

	"github.com/faiq/dopepope/routers"
	"gopkg.in/mgo.v2"
)

func main() {
	if err != nil {
		return nil, err
	}
	uri := "mongodb://localhost/"
	sess, err := mgo.Dial(uri)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer sess.Close()
	r := MakeRouter(sess)
	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

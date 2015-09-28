package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/faiq/dopepope/populate"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type Rhyme struct {
	Word      string `json:"word"`
	Freq      int    `json:"freq"`
	Score     int    `json:"score"`
	Flags     string `json:"flags"`
	Syllables string `json:"syllables"`
}

const url = "http://rhymebrain.com/talk?function=getRhymes&maxResults=30&word="

var fireFlag string

func MakeRequest(mainWait *sync.WaitGroup, term string) ([]string, error) {
	var fire []string
	defer mainWait.Done()
	newUrl := url + term
	resp, err := http.Get(newUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var Rhymes []Rhyme
	err = json.Unmarshal([]byte(jsonDataFromHttp), &Rhymes)
	if err != nil {
		return nil, err
	}
	uri := "mongodb://localhost/"
	sess, err := mgo.Dial(uri)
	if err != nil {
		return nil, err
	}
	defer sess.Close()
	if err != nil {
		return nil, err
	}
	updates := make(chan populate.Sentence)
	var chanWait sync.WaitGroup
	for _, word := range Rhymes {
		time.Sleep(time.Millisecond * 600)
		chanWait.Add(1)
		go RunQuery(word.Word, updates, sess, &chanWait)
	}

	go func() {
		chanWait.Wait()
		close(updates)
	}()
	for result := range updates {
		fire = append(fire, result.populate.Sentence)
	}
	// Wait for all the queries to complete.
	return fire, nil
}

func RunQuery(query string, sendUpdates chan<- populate.Sentence, mongoSession *mgo.Session, waitGroup *sync.WaitGroup) {
	// Decrement the wait group count so the program knows this
	// has been completed once the goroutine exits.

	// Request a socket connection from the session to process our query.
	// Close the session when the goroutine exits and put the connection back
	// into the pool.
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	defer waitGroup.Done()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB("dopepope").C("sentencestest")

	log.Printf("RunQuery : %d : Executing\n", query)
	var result populate.Sentence
	err := collection.Find(bson.M{"lastWord": query}).One(&result)
	time.Sleep(time.Second * 5)
	if err != nil {
		log.Printf("RunQuery : ERROR : %s\n", err)
		return
	}
	sendUpdates <- result

}

func main() {
	flag.StringVar(&fireFlag, "fire", "war", "whatchu want the pope to rap about????????")
	flag.Parse()
	var wait sync.WaitGroup
	wait.Add(1)
	fire, err := MakeRequest(&wait, fireFlag)
	wait.Wait()
	if err != nil {
		fmt.Printf("I hope this doesnt break at the demo............")
	}
	filename := "output.txt"

	file, err := os.Create(filename)

	if err != nil {
		fmt.Println(err)
	}

	for _, fireline := range fire {
		n, err := io.WriteString(file, fireline)
		if err != nil {
			fmt.Println(n, err)
		}
	}
	file.Close()
}

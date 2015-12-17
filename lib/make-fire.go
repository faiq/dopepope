package lib

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/faiq/dopepope/populate"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

type Rhyme struct {
	Word      string `json:"word"`
	Freq      int    `json:"freq"`
	Score     int    `json:"score"`
	Flags     string `json:"flags"`
	Syllables string `json:"syllables"`
}

const url = "http://rhymebrain.com/talk?function=getRhymes&maxResults=50&word="
const MAXCONCURRENCY = 5

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
	throttle := make(chan int, MAXCONCURRENCY)
	var chanWait sync.WaitGroup
	chanWait.Add(len(Rhymes))
	go func() {
		for result := range updates {
			if result.Sentence != "" {
				fire = append(fire, result.Sentence)
			}
		}
	}()
	for _, word := range Rhymes {
		throttle <- 1
		go RunQuery(word.Word, updates, sess, &chanWait, throttle)
	}
	chanWait.Wait()
	close(updates)
	close(throttle)
	return fire, nil
}

func RunQuery(query string, sendUpdates chan<- populate.Sentence, mongoSession *mgo.Session, waitGroup *sync.WaitGroup, throttle chan int) {
	// Decrement the wait group count so the program knows this
	// has been completed once the goroutine exits.

	// Request a socket connection from the session to process our query.
	// Close the session when the goroutine exits and put the connection back
	// into the pool.
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	defer func() {
		<-throttle
		waitGroup.Done()
	}()
	// Get a collection to execute the query against.
	collection := sessionCopy.DB("dopepope").C("sentencestest")
	var result populate.Sentence
	err := collection.Find(bson.M{"lastWord": query}).One(&result)
	if err != nil && err != mgo.ErrNotFound {
		fmt.Printf("%v \n", err)
		<-throttle
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
		fmt.Printf("%v \n", err)
	}
	filename := "output.txt"

	file, err := os.Create(filename)

	if err != nil {
		fmt.Println(err)
	}

	for _, fireline := range fire {
		n, err := io.WriteString(file, fireline+"\n")
		if err != nil {
			fmt.Println(n, err)
		}
	}
	file.Close()
}

package populate

import (
	"bytes"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Sentence struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	LastWord string        `bson:"lastWord"`
	Sentence string        `bson:"sentence"`
}

func readLines(filename string) ([]string, error) {
	var lines []string
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return lines, err
	}
	buf := bytes.NewBuffer(file)
	for {
		line, err := buf.ReadString('\n')
		if len(line) == 0 {
			if err != nil {
				if err == io.EOF {
					break
				}
				return lines, err
			}
		}
		lines = append(lines, line)
		if err != nil && err != io.EOF {
			return lines, err
		}
	}
	return lines, nil
}

func cleanLinesAndSave(lines []string) error {
	uri := "mongodb://localhost/"
	sess, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	defer sess.Close()
	if err != nil {
		return err
	}
	c := sess.DB("dopepope").C("sentences")
	for _, line := range lines {
		lastIndex := strings.LastIndex(line, " ")
		if lastIndex == -1 {
			continue
		}
		lastWord := line[lastIndex+1 : len(line)-2]
		lastWord = strings.ToLower(lastWord)
		if len(lastWord) > 2 && strings.ContainsAny(lastWord, "a b c d e f g h i j k l m n o p q r s t u v w x y z") {
			line = strings.TrimSpace(line)
			err = c.Insert(&Sentence{bson.NewObjectId(), lastWord, line})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	pwd, _ := os.Getwd()
	fileName := pwd + "/speeches/final.txt"
	fileName, _ = filepath.Abs(fileName)
	lines, err := readLines(fileName)
	if err != nil {
		fmt.Println(err)
	}
	err = cleanLinesAndSave(lines)
	if err != nil {
		fmt.Println(err)
	}
}

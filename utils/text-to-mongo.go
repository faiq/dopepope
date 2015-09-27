package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

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

func delimitLines(lines []string) {
	var ret []string
	terminals := []string{"?", ".", "!"}
	lineEnded := false
	lastLine := ""
	fmt.Printf("%d ", len(lines))
LineLoop:
	for i, line := range lines {
		for i, terminal := range terminals {
			// if the last line didnt end a sentence
			// add that on to the lastLine string
			if lineEnded == false && lastLine != "" {
				if term := strings.Index(line, terminal); term > 0 {
					lastLine = strings.Join([]string{lastLine, line[:term+1]}, " ")
					ret = append(ret, lastLine)
					line = line[term+1:]
					lineEnded = true
					continue LineLoop
				} else {
					if i == len(terminals)-1 {
						lastLine = strings.Join([]string{lastLine, line[:term+1]}, " ")
					}
					// try with another delim
				}
			}
			if term := strings.Index(line, terminal); term > 0 {
				ret = append(ret, line[:term+1])
				line = line[term+1:]
				lineEnded = true
				continue LineLoop
			} else {
				// so we didnt find a delimeter this time around if its the last time we should make lastLine, line
				if i == len(terminals)-1 {
					lastLine = line
					lineEnded = false
				}
			}
		}
	}
}

func main() {
	files, err := ioutil.ReadDir("../speeches")
	if err != nil {
		fmt.Printf("err reading dir")
	}
	for _, f := range files {
		fileName := fmt.Sprintf("../speeches/%s", f.Name())
		lines, err := readLines(fileName)
		fmt.Printf("%v", lines)
		if err != nil {
			fmt.Println(err)
		}
		delimitLines(lines)
	}
}

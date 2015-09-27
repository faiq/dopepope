package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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

func cleanLines(lines []string) []string {
	var ret []string
	for _, line := range lines {
		if len(line) > 4 {
			ret = append(ret, line)
		}
	}
	return ret
}

func main() {
	fileName := "../speeches/final.txt"
	lines, err := readLines(fileName)
	if err != nil {
		fmt.Println(err)
	}
	finalLines := cleanLines(lines)
	fmt.Printf("%v", finalLines)
}

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

func main() {
	files, err := ioutil.ReadDir("../speeches")
	if err != nil {
		fmt.Printf("err reading dir")
	}
	for _, f := range files {
		fileName := fmt.Sprintf("../speeches/%s", f.Name())
		lines, err := readLines(fileName)
		fmt.Println(len(lines))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

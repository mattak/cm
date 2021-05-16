package cmd

import (
	"encoding/json"
	"firebase.google.com/go/messaging"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func ReadStdinLines() []string {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	allText := string(bytes)
	allText = strings.TrimRight(allText, "\n")
	lines := strings.Split(allText, "\n")

	return lines
}

func ReadMessageFromStdinJson(lines []string) ([]*messaging.Message, error) {
	results := make([]*messaging.Message, len(lines))

	for i := 0; i < len(lines); i++ {
		msg, err := ParseMessageLine(lines[i])
		if err != nil {
			return nil, err
		}
		results[i] = msg
	}
	return results, nil
}

func ParseMessageLine(line string) (*messaging.Message, error) {
	data := []byte(line)
	var msg messaging.Message
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

func TrimEmptyLines(lines []string) []string {
	newLines := []string{}
	emptyRegex := regexp.MustCompile("^#?\\s*$")

	for i := 0; i < len(lines); i++ {
		if !emptyRegex.MatchString(lines[i]) {
			newLines = append(newLines, lines[i])
		}
	}

	return newLines
}

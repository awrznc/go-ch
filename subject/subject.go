package subject

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Thread struct {
	ThreadKey     string
	Title         string
	ResponseCount uint64
}

type Subject struct {
	Threads []Thread
}

func Deserialize(subject *Subject, target string) error {
	lines := strings.Split(target, "\n")
	regexpString := regexp.MustCompile(`^(.+.dat)<>(.+?)\s+\((\d+)\)$`)

	for _, line := range lines {
		matches := regexpString.FindStringSubmatch(line)

		if len(matches) != 4 {
			return fmt.Errorf("invalid string: %v", line)
		}

		responseCount, err := strconv.ParseUint(matches[3], 10, 64)
		if err != nil {
			return err
		}

		thread := Thread{
			ThreadKey:     matches[1],
			Title:         matches[2],
			ResponseCount: responseCount,
		}

		subject.Threads = append(subject.Threads, thread)
	}

	return nil
}

func Serialize(subject *Subject) string {
	result := []string{}

	for _, information := range subject.Threads {
		result = append(result, information.ThreadKey+`<>`+information.Title+`  (`+strconv.FormatUint(information.ResponseCount, 10)+`)`)
	}

	return strings.Join(result, "\n")
}

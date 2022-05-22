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
		line := Replace(information.ThreadKey+`<>`+information.Title+` (`+strconv.FormatUint(information.ResponseCount, 10)+`)`)
		result = append(result, line)
	}

	return strings.Join(result, "\n")
}

func Replace(target string) string {
	var htmlEscaper = strings.NewReplacer(
		`&`, "&amp;",
		`'`, "&apos;",
		`<`, "&lt;",
		`>`, "&gt;",
		`"`, "&quot;",
		"\n", "<br>",
		"\r", "",
	)

	return " " + htmlEscaper.Replace(target) + " "
}

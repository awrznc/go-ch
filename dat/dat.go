package dat

import (
	"fmt"
	"regexp"
	"strings"
)

type Dat struct {
	Responses []Response
}

type Response struct {
	Name        string
	MailAddress string
	Datetime    string
	Id          string
	Body        string
	Title       string
}

func Deserialize(dat *Dat, target string) error {
	lines := strings.Split(target, "\n")
	regexpString := regexp.MustCompile(`^(.+?)<>(.+?)<>(.+?)\sID:(.+?)<>(.+?)<>(.*)$`)

	for _, line := range lines {
		matches := regexpString.FindStringSubmatch(line)

		if len(matches) != 7 {
			return fmt.Errorf("invalid string: %v", line)
		}

		response := Response{
			Name:        matches[1],
			MailAddress: matches[2],
			Datetime:    matches[3],
			Id:          matches[4],
			Body:        matches[5],
			Title:       matches[6],
		}

		dat.Responses = append(dat.Responses, response)
	}

	return nil
}

func Serialize(dat *Dat) string {
	result := []string{}

	for _, response := range dat.Responses {
		result = append(result, response.Name+`<>`+response.MailAddress+`<>`+response.Datetime+` ID:`+response.Id+`<>`+response.Body+`<>`+response.Title)
	}

	return strings.Join(result, "\n")
}

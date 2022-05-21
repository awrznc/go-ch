package handler

import (
	"fmt"
	"net/http"
	"regexp"
	// "time"
	"io/ioutil"
    "strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	gh "ch/examples/5kan/github"
	sb "ch/subject"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	subject := sb.Subject{Threads: make([]sb.Thread, 0)}

	matches := regexp.MustCompile(`^/(.+?)/(subject.txt)$`).FindStringSubmatch(request.URL.Path)
	id := matches[1]

	var repo gh.Repo
	gh.GetRepo(&repo, id)

	for i := 1; i <= (repo.OpenIssues/100)+1; i++ {
		var issues []gh.Issue
		gh.GetIssues(&issues, id, i)

		for _, issue := range issues {
			// datetime, err := time.Parse("2006-01-02T15:04:05Z", issue.CreatedAt)
			// if err != nil {
			// 	panic(err)
			// }

			subject.Threads = append(
				subject.Threads,
				sb.Thread{
					// ThreadKey:     fmt.Sprintf("%v.dat", datetime.Unix()),
					ThreadKey:     fmt.Sprintf("%v.dat", issue.Number),
					Title:         issue.Title,
					ResponseCount: issue.Comments + 1,
				},
			)
		}
	}

	sjis := utf8ToSjis(sb.Serialize(&subject)+"\n")

	writer.Write([]byte(sjis))
}

func utf8ToSjis(target string) string {
    stringReader := strings.NewReader(target)
    transformReader := transform.NewReader(stringReader, japanese.ShiftJIS.NewEncoder())
    result, err := ioutil.ReadAll(transformReader)
    if err != nil {
        panic(err)
    }
    return string(result)
}

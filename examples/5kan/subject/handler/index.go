package handler

import (
	"fmt"
	"net/http"
	"regexp"

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
			subject.Threads = append(
				subject.Threads,
				sb.Thread{
					ThreadKey:     fmt.Sprintf("%v.dat", issue.Number),
					Title:         issue.Title,
					ResponseCount: issue.Comments + 1,
				},
			)
		}
	}

	writer.Write([]byte(sb.Serialize(&subject)))
}

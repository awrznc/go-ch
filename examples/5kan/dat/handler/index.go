package handler

import (
	"net/http"
	"regexp"
	"strings"
	"fmt"
	"time"

	dt "ch/dat"
	gh "ch/examples/5kan/github"
)

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

func getIssueNumber(repoId string, datId string) string {
	var repo gh.Repo
	gh.GetRepo(&repo, repoId)

	for i := 1; i <= (repo.OpenIssues/100)+1; i++ {
		var issues []gh.Issue
		gh.GetIssues(&issues, repoId, i)

		for _, issue := range issues {
			datetime, err := time.Parse("2006-01-02T15:04:05Z", issue.CreatedAt)
			if err != nil {
				panic(err)
			}
			unixtime := datetime.Unix()
			if fmt.Sprintf("%v", unixtime) == datId {
				return fmt.Sprintf("%v", issue.Number)
			}
		}
	}
	panic(fmt.Errorf("not found."))
}

func Handler(writer http.ResponseWriter, request *http.Request) {
	dat := dt.Dat{Responses: make([]dt.Response, 0)}

	matches := regexp.MustCompile(`^/(.+?)/(.+?)/(.+?).dat$`).FindStringSubmatch(request.URL.Path)
	repoId := matches[1]
	datId := matches[3]

	issueNumber := getIssueNumber(repoId, datId)

	var issue gh.Issue
	gh.GetIssue(&issue, repoId, issueNumber)

	iBody := Replace(issue.Body)

	dat.Responses = append(
		dat.Responses,
		dt.Response{
			Name        : issue.User.Login,
			MailAddress : "",
			Datetime    : issue.UpdatedAt,
			Id          : issue.User.Id,
			Body        : iBody,
			Title       : issue.Title,
		},
	)

	for i := 1; i <= (int(issue.Comments)/100)+1; i++ {
		var comments []gh.Comment
		gh.GetComments(&comments, repoId, issueNumber, i)

		for _, comment := range comments {
			cBody := Replace(comment.Body)

			dat.Responses = append(
				dat.Responses,
				dt.Response{
					Name        : comment.User.Login,
					MailAddress : "",
					Datetime    : comment.UpdatedAt,
					Id          : comment.User.Id,
					Body        : cBody,
					Title       : "",
				},
			)
		}
	}

	writer.Write([]byte(dt.Serialize(&dat)))
}

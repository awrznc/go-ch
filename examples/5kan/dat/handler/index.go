package handler

import (
	"net/http"
	"regexp"
	"strings"

	dt "ch/dat"
	gh "ch/examples/5kan/github"
)

func Replace(target string) string {
	result := strings.Replace(target, "<", "&lt;", -1)
	result = strings.Replace(result, ">", "&gt", -1)
	result = strings.Replace(result, "&", "&amp;", -1)
	result = strings.Replace(result, `"`, "&quot;", -1)
	result = strings.Replace(result, "\n", "<br>", -1)
	result = strings.Replace(result, "\r", "", -1)

	return " " + result + " "
}

func Handler(writer http.ResponseWriter, request *http.Request) {
	dat := dt.Dat{Responses: make([]dt.Response, 0)}

	matches := regexp.MustCompile(`^/(.+?)/(.+?).dat$`).FindStringSubmatch(request.URL.Path)
	repoId := matches[1]
	issueNumber := matches[2]

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

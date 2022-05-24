package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"time"
	"io/ioutil"
    "strings"
	"compress/gzip"
	"bytes"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"github.com/rivo/uniseg"

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
			datetime, err := time.Parse("2006-01-02T15:04:05Z", issue.CreatedAt)
			if err != nil {
				panic(err)
			}

			subject.Threads = append(
				subject.Threads,
				sb.Thread{
					ThreadKey:     fmt.Sprintf("%v.dat", datetime.Unix()),
					// ThreadKey:     fmt.Sprintf("%v.dat", issue.Number),
					Title:         issue.Title,
					ResponseCount: issue.Comments + 1,
				},
			)
		}
	}

	sjis := utf8ToSjis(sb.Serialize(&subject)+"\n")

	var result string
	writer.Header().Set("Content-Type", "text/plain; charset=Shift_JIS")

	if strings.Contains(request.Header.Get("Accept-Encoding"), "gzip") {
		result = gzipping(sjis)
		writer.Header().Set("Content-Encoding", "gzip")
	} else {
		result = sjis
	}

	writer.Write([]byte(result))
}

func gzipping(target string) string {
	var buffer bytes.Buffer
	writer := gzip.NewWriter(&buffer)
	writer.Write([]byte(target))
	writer.Close()
	return string(buffer.Bytes())
}

func utf8ToSjis(target string) string {

	// get encord target
	replaceChars := make([]string, 0)
	gr := uniseg.NewGraphemes(target)
    for gr.Next() {
        rs := gr.Runes()
		for _, emoji := range rs {

			// unless ShiftJIS
			if emoji >= 0xEA5C {
				replaceChars = append(replaceChars, string(emoji), fmt.Sprintf("&#%v;", emoji))
			}
		}
    }

	replacedTarget := strings.NewReplacer(replaceChars...).Replace(target)

	stringReader := strings.NewReader(replacedTarget)
    transformReader := transform.NewReader(stringReader, japanese.ShiftJIS.NewEncoder())
    result, err := ioutil.ReadAll(transformReader)
    if err != nil {
        panic(err)
    }
    return string(result)
}

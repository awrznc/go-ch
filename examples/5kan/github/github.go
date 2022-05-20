package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}

func GetUser(user *User, userName string) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.github.com/users/%v", userName), nil)
	if err != nil {
		panic(err)
	}
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(response.Body)

	json.Unmarshal(body, user)
}

type Repo struct {
	Id         uint64 `json:"id"`
	Name       string `json:"name"`
	OpenIssues int    `json:"open_issues"`
}

func GetRepo(repo *Repo, id string) {

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.github.com/repositories/%v", id), nil)
	if err != nil {
		panic(err)
	}
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(response.Body)

	json.Unmarshal(body, repo)
}

func GetRepos(repos *[]Repo, userName string, page int) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.github.com/users/%v/repos?per_page=100&page=%v", userName, page), nil)
	if err != nil {
		panic(err)
	}
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(response.Body)

	json.Unmarshal(body, repos)
}

type Issue struct {
	Number    uint64   `json:"number"`
	Title     string   `json:"title"`
	Comments  uint64   `json:"comments"`
	Body      string   `json:"body"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	User      UserInfo `json:"user"`
}

func GetIssue(issue *Issue, id string, number string) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.github.com/repositories/%v/issues/%v", id, number), nil)
	if err != nil {
		panic(err)
	}
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(response.Body)

	json.Unmarshal(body, issue)
}

func GetIssues(issues *[]Issue, id string, page int) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.github.com/repositories/%v/issues?per_page=100&page=%v", id, page), nil)
	if err != nil {
		panic(err)
	}
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(response.Body)

	json.Unmarshal(body, issues)
}

type UserInfo struct {
	Id    string `json:"id"`
	Login string `json:"login"`
}

type Comment struct {
	Body      string   `json:"body"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	User      UserInfo `json:"user"`
}

func GetComments(comments *[]Comment, id string, number string, page int) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.github.com/repositories/%v/issues/%v/comments?per_page=100&page=%v", id, number, page), nil)
	if err != nil {
		panic(err)
	}
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(response.Body)

	json.Unmarshal(body, comments)
}

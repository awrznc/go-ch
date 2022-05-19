package github

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

type User struct {
	Id uint64 `json:"id"`
	Name string `json:"name"`
	PublicRepos int `json:"public_repos"`
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
	Id uint64 `json:"id"`
	Name string `json:"name"`
	OpenIssues uint64 `json:"open_issues"`
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

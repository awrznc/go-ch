package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	bm "ch/bbsmenu"
	gh "ch/examples/5kan/github"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	dateTime := time.Now()
	unixTime := dateTime.Unix()

	list := bm.MenuList{
		CategoryContent: make([]bm.CategoryContent, 0),
		CategoryTotal:   1,
		CategoryName:    "GitHub",
		CategoryNumber:  "1",
	}

	menu := bm.BBSmenu{
		Description:      "5kan bbsmenu",
		LastModify:       unixTime,
		LastModifyString: dateTime.Format("2006/01/02(Mon) 15:04:05"),
		MenuList:         make([]bm.MenuList, 0),
	}

	menu.MenuList = append(menu.MenuList, list)

	var user gh.User
	gh.GetUser(&user, "github")

	for i := 1; i <= (user.PublicRepos/100)+1; i++ {
		var repos []gh.Repo
		gh.GetRepos(&repos, "github", i)

		for _, repo := range repos {
			menu.MenuList[0].CategoryContent = append(
				menu.MenuList[0].CategoryContent,
				bm.CategoryContent{
					BoardName:     repo.Name,
					Category:      1,
					CategoryName:  "GitHub",
					CategoryOrder: 1,
					DirectoryName: fmt.Sprintf("%v", repo.Id),
					Url:           fmt.Sprintf("https://%v/%v/", request.Host, repo.Id),
				},
			)
		}
	}

	menuString, err := json.Marshal(menu)
	if err != nil {
		panic(err)
	}
	writer.Write([]byte(menuString))
}

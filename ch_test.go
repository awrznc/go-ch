package ch

import (
	"encoding/json"
	"testing"
	"time"

	bbsmenu "ch/bbsmenu"
	dat "ch/dat"
	subject "ch/subject"
)

func TestBBSmenu(tester *testing.T) {
	dateTime := time.Now()
	unixTime := dateTime.Unix()
	content := bbsmenu.CategoryContent{
		BoardName:     "github",
		Category:      1,
		CategoryName:  "GitHub",
		CategoryOrder: 1,
		DirectoryName: "github",
		Url:           "https://5kan.5kan/github/",
	}

	list := bbsmenu.MenuList{
		CategoryContent: []bbsmenu.CategoryContent{content},
		CategoryTotal:   1,
		CategoryName:    "GitHub",
		CategoryNumber:  "1",
	}

	menu := bbsmenu.BBSmenu{
		Description:      "5kan bbsmenu",
		LastModify:       unixTime,
		LastModifyString: dateTime.Format("2006/01/02(Mon) 15:04:05"),
		MenuList:         []bbsmenu.MenuList{list},
	}

	_, err := json.Marshal(menu)
	if err != nil {
		tester.Errorf("failed to Marshal for BBSmenu: %v", err)
	}
}

func TestSubject(tester *testing.T) {
	target := `1.dat<>スレ1  (1)
20.dat<>スレ2  (10)
300.dat<>スレ3  (100)`
	object := subject.Subject{}
	err := subject.Deserialize(&object, target)
	if err != nil {
		tester.Errorf("failed to Deserialize for Subject: %v", err)
	}
	if len(object.Threads) != 3 {
		tester.Errorf("subject size: %v", len(object.Threads))
	}

	result := subject.Serialize(&object)
	if target != result {
		tester.Errorf("failed to Serialize for Subject: \n\n%v\n\n%v", target, result)
	}
}

func TestDat(tester *testing.T) {
	target := `nanashisan<>5kan@5kan.5kan<>2000/01/01(土) 00:00:00.00 ID:shobon<> (´・ω・｀) <>テスト
nanashisan<>5kan@5kan.5kan<>2000/01/01(土) 00:00:01.00 ID:otsu<> (　´∀｀) <>
nanashisan<>5kan@5kan.5kan<>2000/01/01(土) 00:00:02.00 ID:iine<> (・∀・) <>`
	object := dat.Dat{}
	err := dat.Deserialize(&object, target)
	if err != nil {
		tester.Errorf("failed to Deserialize for Dat: %v", err)
	}
	if len(object.Responses) != 3 {
		tester.Errorf("dat size: %v", len(object.Responses))
	}

	result := dat.Serialize(&object)
	if target != result {
		tester.Errorf("failed to Serialize for Dat: \n\n%v\n\n%v", target, result)
	}
}

package main

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/gogf/gf/os/gfile"
)

type Question2Answer struct {
	WorldName    string `json:"worldName" csv:"worldName"`
	RightAnswer  string `json:"rightAnswer" csv:"rightAnswer"`
	WrongAnswer1 string `json:"wrongAnswer1" csv:"wrongAnswer1"`
	WrongAnswer2 string `json:"wrongAnswer2" csv:"wrongAnswer2"`
	WrongAnswer3 string `json:"wrongAnswer3" csv:"wrongAnswer3"`
	Dir          string `json:"dir" csv:"dir"`
	Img          string `json:"img" csv:"img"`
}

type LevelData struct {
	Dir  string `json:"dir"`
	List []struct {
		Options []string `json:"options"`
		Image   string   `json:"image"`
		Dir     string   `json:"dir"`
	} `json:"list"`
}

func main() {
	f, err := os.OpenFile("answer.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var questions []*Question2Answer

	if err := gocsv.UnmarshalFile(f, &questions); err != nil {
		panic(err)
	}
	var levelDatas []LevelData
	for i := 0; i < len(questions); i = i + 4 {
		var item LevelData
		item.Dir = strings.TrimSpace(questions[i].WorldName)
		for j := i; j < i+4; j++ {
			var options []string
			options = append(options, strings.TrimSpace(questions[j].RightAnswer))
			options = append(options, strings.TrimSpace(questions[j].WrongAnswer1))
			options = append(options, strings.TrimSpace(questions[j].WrongAnswer2))
			options = append(options, strings.TrimSpace(questions[j].WrongAnswer3))
			round := struct {
				Options []string `json:"options"`
				Image   string   `json:"image"`
				Dir     string   `json:"dir"`
			}{
				Options: options,
				Image:   questions[j].Img,
				Dir:     questions[j].Dir,
			}
			item.List = append(item.List, round)
		}
		levelDatas = append(levelDatas, item)
	}

	contentByte, _ := json.Marshal(levelDatas)

	outPath := "level.js"
	_ = gfile.PutContents(outPath, "let levelData = ")
	_ = gfile.PutBytesAppend(outPath, contentByte)
}

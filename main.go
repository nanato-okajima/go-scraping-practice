package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"log"

	"go_webscraping_practice/scraping"
)

const (
	indexUrl = "https://www.douwa-douyou.jp/contents/html/douwa/douwa6.shtml"
	category = "日本の昔話のタイトル一覧"
	title    = "あかんぼうをたべた女"
)

func main() {
	doc, err := scraping.GetDocument(indexUrl)
	if err != nil {
		log.Fatal(err)
	}

	storyList := scraping.FetchTitles(doc, category)
	doc, err = scraping.GetDocument(storyList[title])
	if err != nil {
		log.Fatal(err)
	}

	story := scraping.GetFairyTale(doc)
	s, err := json.Marshal(story)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(s))
}

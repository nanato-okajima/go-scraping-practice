package main

import "go_webscraping_practice/scraping"

// "encoding/json"

const (
	indexUrl  = "https://www.douwa-douyou.jp/contents/html/douwa/douwa6.shtml"
	nikkeiUrl = "https://tech.nikkeibp.co.jp"
	category  = "日本の昔話のタイトル一覧"
	title     = "あかんぼうをたべた女"
)

func main() {
	scraping.Attendance()
	// doc, err := scraping.GetDocument(nikkeiUrl)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// scraping.GetHoge(doc)

	// storyList := scraping.FetchTitles(doc, category)
	// doc, err = scraping.GetDocument(storyList[title])
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// story := scraping.GetFairyTale(doc)
	// s, err := json.Marshal(story)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(s))
}

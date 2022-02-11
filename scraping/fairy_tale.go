package scraping

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Story struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

const (
	domain    = "https://www.douwa-douyou.jp"
	storyPath = "contents/html"
)

func FetchTitles(doc *goquery.Document, category string) map[string]string {
	var url string
	titleLink := doc.Find(fmt.Sprintf("table[summary='%s']", category)).Find("td > a")
	storyPages := make(map[string]string, len(titleLink.Nodes))
	titleLink.Each(func(i int, s *goquery.Selection) {
		p, ok := s.Attr("href")
		if ok == true {
			url = toStoryPageURL(p)
		}
		storyPages[s.Text()] = url
	})
	return storyPages
}

func GetFairyTale(doc *goquery.Document) Story {
	title := doc.Find("h4").Text()
	content := doc.Find("p.story").Text()

	return Story{
		title,
		content,
	}
}

func toStoryPageURL(path string) string {
	cp := strings.TrimLeft(path, "../")
	urlParts := []string{domain, storyPath, cp}
	url := strings.Join(urlParts, "/")

	return url
}

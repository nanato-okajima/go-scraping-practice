package scraping

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

func GetDocument(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	det := chardet.NewTextDetector()
	detRslt, err := det.DetectBest(buf)
	if err != nil {
		log.Println(err)
	}

	bReader := bytes.NewReader(buf)
	reader, err := charset.NewReaderLabel(detRslt.Charset, bReader)
	if err != nil {
		log.Println(err)
	}

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Println(err)
	}

	return doc, nil
}

package site

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strings"
)

type News struct {
	Num     string
	Title   string
	Text    string
	Photos  []string
	Hashtag string
	Links   []string
}

const (
	div               = "div"
	a                 = "a"
	p                 = "p"
	href              = "href"
	pageNews          = "/p/news/"
	blockAllNewsClass = ".listNews"
	blockNewsClass    = ".oneArticle"
	blockNewsDate     = ".date"
	blockNewsText     = ".textArticle"
)

var (
	url = "http://school46.irk.ru"
)

func Init() {
	res, err := connect(url + pageNews)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := getHtmlDocument(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	getListNews(doc)
}

func connect(urlNews string) (*http.Response, error) {
	// Request the HTML page.
	res, err := http.Get(urlNews)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return res, nil
}

func getHtmlDocument(htmlBody io.Reader) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(htmlBody)
}

func getListNews(htmlDocument *goquery.Document) {
	htmlDocument.Find(blockNewsClass).Each(func(i int, s *goquery.Selection) {
		getNews(s)
	})
}

func getNews(selection *goquery.Selection) {
	date := getNewsDate(selection)
	blockNews := selection.Find(blockNewsText)
	title := getNewsTitle(blockNews)

	fmt.Printf("Review: %s %s\n", date, title)
}

func getNewsDate(s *goquery.Selection) string {
	return s.Find(blockNewsDate).Text()
}

func getNewsTitle(s *goquery.Selection) string {
	var title []string
	s.Find("a").Each(func(i int, selection *goquery.Selection) {
		a := strings.TrimSpace(selection.Text())
		if a != "" {
			title = append(title, a)
		}
	})
	return title[0]
}

func getNewsText(s *goquery.Selection) string {
	return ""
}

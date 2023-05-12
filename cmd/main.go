package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := getUrl()
	crawlerInit(url)
}

func getUrl() string {
	fmt.Print("Digite a URL (Ex: https://www.sua-url.com): ")
	var url string
	fmt.Scanln(&url)

	checkUrl(url)

	return url
}

func checkUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(err)
	}
}

func crawlerInit(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			if !strings.HasPrefix(href, "#") && !strings.HasPrefix(href, "mailto:") {
				fmt.Println(href)
			}
		}
	})
}

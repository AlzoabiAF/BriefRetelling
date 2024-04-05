package scrapHTML

import (
	"github.com/gocolly/colly"
	"strings"
)

type Data struct {
	Heading  string
	TextItem []string
}

func Scrap(url string) *Data {
	c := colly.NewCollector()
	var data Data
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		data.Heading = e.Text
	})

	c.OnHTML("ul", func(e *colly.HTMLElement) {
		data.TextItem = strings.Split(e.Text, "•")
	})

	c.Visit(url)
	return &data
	// TODO: logger
}

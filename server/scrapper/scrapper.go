package scrapper

import (
	"fmt"

	"github.com/gocolly/colly"
)

func ScrappingWIKI() {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	// Find and print all links
	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")
		fmt.Println(links)
	})
	c.Visit("https://en.wikipedia.org/wiki/Web_scraping")
}

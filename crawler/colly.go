package crawler

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly"
)

//Crawler
func Crawler(target string) {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.URLFilters(
			regexp.MustCompile(".*" + target + "*"),
		),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on target
	c.Visit("http://" + target)
}

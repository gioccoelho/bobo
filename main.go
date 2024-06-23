package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	// Declaring the url and initialising the collector
	url := "https://github.com/trending"
	collector := colly.NewCollector()

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Connecting to", r.URL)
	})
	// Handling the callbacks
	collector.OnResponse(func(r *colly.Response) {
		fmt.Printf("Connected to %v\n\n", r.Request.URL)
	})
	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("Ooops, an error ocurred:", e)
	})

	collector.OnHTML("article.Box-row", func(e *colly.HTMLElement) {
		// Extract repository name and URL
		repo_name := e.ChildText("h2.h3.lh-condensed")

		// Removing all new lines and blank spaces
		repo_name = strings.ReplaceAll(repo_name, " ", "")
		repo_name = strings.ReplaceAll(repo_name, "\n", "")

		// Catching additional info (stars and forks)
		additional_info := e.ChildText("div.f6.color-fg-muted.mt-2 a")

		// Removing all new lines and blank spaces
		additional_info = strings.ReplaceAll(additional_info, " ", "")
		additional_info = strings.ReplaceAll(additional_info, "\n", "")

		// Extracting stars and forks from additional_info
		stars := additional_info[:strings.Index(additional_info, ",")+4]
		forks := additional_info[strings.Index(additional_info, ",")+4:]

		// Extract the repo description
		description := e.ChildText("p.col-9.color-fg-muted.my-1.pr-4")

		fmt.Printf("Repository: %s --- Stars: %s --- Forks: %s\nDescription: %s\n\n", repo_name, stars, forks, description)
	})

	collector.Visit(url)
}

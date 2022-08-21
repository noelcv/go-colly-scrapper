package main

import (
	"fmt"
	"log"
	"github.com/playwright-community/playwright-go"
)

type news struct {
	Headline string `json="title"`
	Source string `json="source"`
}

func main() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://theguardian.com"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	entries, err := page.QuerySelectorAll(".fc-item__container")
	if err != nil {
		log.Fatalf("could not get entries: %v", err)
	}
	
	
	var newsSlice []news
	for i, entry := range entries {
		if i < 5 {
			
			titleElement, err := entry.QuerySelector("a[data-link-name=article]")
			if err != nil {
				log.Fatalf("could not get title element: %v", err)
			}
			title, err := titleElement.TextContent()
			if err != nil {
				log.Fatalf("could not get text content: %v", err)
			}
			
			url, err := titleElement.GetAttribute("href")
			if err != nil {
				log.Fatalf("could not get text content: %v", err)
			}
			
			news := news {
				Headline: title,
				Source: url,
			}
			
			newsSlice = append(newsSlice, news)
			// fmt.Printf("%d: %s\n", i+1, title)
		}
	}
	
	fmt.Println(newsSlice)
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}

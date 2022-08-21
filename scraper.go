// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/gocolly/colly"
// )

// type item struct {
// 	Name string `json:"name"`
// 	Price   string `json:"price"`
// 	ImgUrl string `json:"imgurl"`
// }

// func main() {
// 	//instance for Colly Collector
// 	c := colly.NewCollector(
// 		colly.AllowedDomains("j2store.net"),
// 	)
	
// 	var items []item
	
// 	//specify the css selector to scrape, store the Child results in your struct and append them to a slice
// 	c.OnHTML("div.col-sm-4 div[itemprop=itemListElement]", func(h *colly.HTMLElement) {
// 		item := item {
// 			Name: h.ChildText("h2.product-title"),
// 			Price: h.ChildText("div.sale-price"),
// 			ImgUrl: h.ChildAttr("img", "src"),
// 		}
		
// 		items = append(items, item)

// 	})
	
// 	//iterate through following pages
// 	c.OnHTML("[title=Next]", func(h *colly.HTMLElement){
// 		next_page := h.Request.AbsoluteURL(h.Attr("href"))
// 		c.Visit(next_page)
// 	})
	
// 	//print each visited page
// 	c.OnRequest(func(r *colly.Request){
// 		fmt.Println(r.URL.String())
// 	})
	
// 	c.Visit("https://j2store.net/demo/index.php/shop")
	
// 	c.OnError(func(_ *colly.Response, err error) {
//     log.Println("Something went wrong:", err)
// })

// content, err := json.Marshal(items)

// 	if err != nil {
// 	fmt.Println(err.Error())	
// 	}
	
// 	//Save file with UNIX permissions
// 	os.WriteFile("products.json", content, 0644)

// }

package main

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

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
	for i, entry := range entries {
		titleElement, err := entry.QuerySelector("a[data-link-name=article]")
		if err != nil {
			log.Fatalf("could not get title element: %v", err)
		}
		title, err := titleElement.TextContent()
		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}
		fmt.Printf("%d: %s\n", i+1, title)
	}
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}

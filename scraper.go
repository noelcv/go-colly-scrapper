package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type item struct {
	Name string `json:"name"`
	Price   string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {
	//instance for Colly Collector
	c := colly.NewCollector(
		colly.AllowedDomains("j2store.net"),
	)
	
	var items []item
	
	//specify the css selector to scrape, store the Child results in your struct and append them to a slice
	c.OnHTML("div.col-sm-4 div[itemprop=itemListElement]", func(h *colly.HTMLElement) {
		item := item {
			Name: h.ChildText("h2.product-title"),
			Price: h.ChildText("div.sale-price"),
			ImgUrl: h.ChildAttr("img", "src"),
		}
		
		items = append(items, item)

	})
	
	//iterate through following pages
	c.OnHTML("[title=Next]", func(h *colly.HTMLElement){
		next_page := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(next_page)
	})
	
	//print each visited page
	c.OnRequest(func(r *colly.Request){
		fmt.Println(r.URL.String())
	})
	
	c.Visit("https://j2store.net/demo/index.php/shop")
	
	c.OnError(func(_ *colly.Response, err error) {
    log.Println("Something went wrong:", err)
})

content, err := json.Marshal(items)

	if err != nil {
	fmt.Println(err.Error())	
	}
	
	//Save file with UNIX permissions
	os.WriteFile("products.json", content, 0644)

}



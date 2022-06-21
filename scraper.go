package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

// Letting Go know the data types we want
 type item struct {
    Name string `name`
	Price string `price`
	ImgUrl string `imgurl`
 } 

  func main () {
	c := colly.NewCollector(
		colly.AllowedDomains(""),
	)

	var items []item

	// HTML classes of specific info will go in the parenthesis
	c.onHTML("", func(h *colly.HTMLElement) {
		 item := item{
			Name: h.ChildText(""),
			Price: h.ChildText(""),
			ImgUrl: h.ChildAttr(""), 
		 }

		 items = append(items, item)  
	})

	// Scraping through the next page
	c.OnHTML("[], func(h *colly.HTMLElement") {
		next_page := h.Request.AbsoluteURL(h.Attr(""))
		c.Visit(items )
	}

	// Making a New request with Colly
	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String)
	})

	// URL of the site you want to scrape goes in these parenthesis
	c.Visit("") 
	
	content, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err.Error())
	}
	os.WriteFile("", content, 0644)
  }
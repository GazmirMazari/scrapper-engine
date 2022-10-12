package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type listing struct {
   Address  string
   Image    string
   Url      string
   Discount string
}

func main() {
   c := colly.NewCollector()
   c.SetRequestTimeout((120 * time.Second))

   //callbacks
   c.OnHTML()
   
   c.OnRequest(func(r*colly.Request){
	fmt.Println("Visiting", r.URL)
   })

   c.OnResponse(func(r *colly.Response) {
      fmt.Println("Got a response from", r.Request.URL)
  })

  c.OnError(func(r *colly.Response, e error) {
      fmt.Println("Got this error:", e)
  })
   c.Visit("https://www.zillow.com/search/GetSearchPageState.htm?")
}
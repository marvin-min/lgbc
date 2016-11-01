package main

import (
//"go-web"
)
import (
	"fmt"
	"log"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	ExampleScrape()
}
func ExampleScrape() {
	doc, err := goquery.NewDocument("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}


type Dog struct {
}

type Duck struct {

}

func (dog Dog)can()   {
	fmt.Println("Dog can catch mouse!")
}
func (duck Duck)can()   {
	fmt.Print("Duck can catch fish!")
}
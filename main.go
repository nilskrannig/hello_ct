package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

const feedLimit = 5

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.heise.de/rss/heise-atom.xml")

	for i := 0; i < feedLimit; i++ {
		fmt.Println(feed.Items[i].Title)
		fmt.Println(feed.Items[i].Link + "\n")
	}
}

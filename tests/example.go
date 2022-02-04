package main

import (
	"fmt"
	cryptopanic "github.com/jerskisnow/CryptoPanic-GO"
	"github.com/jerskisnow/CryptoPanic-GO/posts"
)

func main() {
	client := cryptopanic.CreateClient("<YOUR_TOKEN_HERE>")

	uiFilter := "news"
	currencies := "BTC,ETH"

	res := client.Posts().Latest(10, posts.Filter{
		Public:     false,
		UI:         &uiFilter,
		Currencies: &currencies,
		//Regions:    nil,
		//Kind:       nil,
	})

	fmt.Println(res[0].Title)
}

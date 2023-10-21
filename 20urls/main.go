package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "www.domain.com/women-shoes?type=high-heels&color=blue"

func main() {
	fmt.Println("handling url in golang")
	fmt.Println(myUrl)

	// parsing
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Print("typr of query params: %T\n", qparams)
	fmt.Println(qparams["index"])

	for _, val := range qparams {
		fmt.Println("param is: ", val)
	}

	// partsOfUrl := &url.URL{
	// 	Scheme: "https",
	// }
}

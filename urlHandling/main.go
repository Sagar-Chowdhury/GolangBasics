package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?cousename=reactjs&paymentid=ghbj5678u"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result.Scheme)
		fmt.Println(result.Host)
		fmt.Println(result.Path)
		fmt.Println(result.Port())
		fmt.Println(result.RawQuery)

	}

	partsOfUrl := &url.URL{

		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}

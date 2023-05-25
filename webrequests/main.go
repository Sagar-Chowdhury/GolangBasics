package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	fmt.Println("Web request sample")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response if of Type: %T\n", response)

	defer response.Body.Close() //duty of the caller to close this.

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(databytes))

}

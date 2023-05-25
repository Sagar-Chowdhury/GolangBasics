package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println(" Welcome to GET and POST operations on server")
	performGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

func performGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close() //it's our duty to close the response.

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, err := ioutil.ReadAll(response.Body) //reads the respons in form of bytes
	if err != nil {
		panic(err)
	}
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
	{
		"cousename":"Let's go with golang",
		"price":0,
		"platform":"learncodeOnline.in"
	}
	       
`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {

	const myurl = "http://localhost:8000/postform"
	data := url.Values{}
	data.Add("Name", "Messi")
	data.Add("Position", "CAM")
	data.Add("Club", "FC Barcelona")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	//https
	fmt.Println(result.Host)
	//lco.dev:3000
	fmt.Println(result.Path)
	// ./learn
	fmt.Println(result.Port())
	// 3000
	fmt.Println(result.RawQuery)
	// coursename=reactjs&paymentid=ghbj456ghb

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	// The type of query params are: url.Values

	fmt.Println(qparams["coursename"])
	// [reactjs]

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	// 	Param is:  [reactjs]
	// Param is:  [ghbj456ghb]

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
	// https://lco.dev/tutcss

}

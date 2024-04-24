package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")
	myURL := "https://jsonplaceholder.typicode.com/todos/1"
	fmt.Printf("Type of url is %T\n", myURL)

	parsedURL, err := url.Parse(myURL)
	if err!= nil{
		fmt.Println("Cant parse URL", err)
		return
	}

	// fmt.Println("Type of URL: %T\n",parsedURL)

	fmt.Println("Scheme: ",parsedURL.Scheme)
	fmt.Println("Host: ",parsedURL.Host)
	fmt.Println("Path: ",parsedURL.Path)
	fmt.Println("RawQuery: ",parsedURL.RawQuery)

	//modify URL components
	parsedURL.Path="/newPath"
	parsedURL.RawQuery="username=iamprince"

	//constructing a URL string from a URL object
	newUrl := parsedURL.String()
	fmt.Println("new Url: ",newUrl)
}
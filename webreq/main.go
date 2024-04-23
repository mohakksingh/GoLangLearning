package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning a web server in GoLang")
	res,err :=http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err!=nil{
		fmt.Println("Erro getting GET Response", err)
		return 
	}
	defer res.Body.Close()
	fmt.Printf("Type of response if %T\n",res)
	// fmt.Println("response is: ",res)

	//Read the response body
	data,err :=ioutil.ReadAll(res.Body)
	if err != nil{
		fmt.Println("Error while reading response",err)
		return
	}
	fmt.Println("response is : ",string(data))

}
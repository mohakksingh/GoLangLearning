package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct{
	UserId int `json:"userID"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func perfomGetRequest(){
	fmt.Println("get operation in GoLang")
	res,err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err!=nil{
		fmt.Println("Error getting: ",err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK{
		fmt.Println("Error in getting response: ",res.Status)
		return
	}

	// data,err := ioutil.ReadAll(res.Body)
	// if err!=nil{
	// 	fmt.Println("Error reading ",err)
	// 	return
	// }
	// fmt.Println("Data is: ",string(data))

	var todo Todo
	err=json.NewDecoder(res.Body).Decode(&todo)
	if err!=nil{
		fmt.Println("Error decoding ",err)
		return
	}
	fmt.Println("Todo: ",todo)
	
	fmt.Println("Title reponse:",todo.Title)
	fmt.Println("Completed Response:",todo.Completed)
}

func performPostRequest(){
	fmt.Println("Post operation in GoLang")
	todo := Todo{
		UserId:23,
		Title:"Mohak Singh",
		Completed: true,
	}

	//convert struct to json
	jsonData,err := json.Marshal(todo)
	if err!=nil{
		fmt.Println("Error marshalling",err)
		return
	}

	//convert json to string
	jsonString :=string(jsonData)

	//convert string data to reader
	jsonReader :=strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos"
	
	//send POST req
	res,err :=http.Post(myUrl,"application/json",jsonReader)
	if err!=nil{
		fmt.Println("Error in post request",err)
		return
	}
	defer res.Body.Close()

	//data, _  :=ioutil.ReadAll(res.Body)
	//fmt.Println("Response: ",string(data))

	fmt.Println("Response status: ",res.Status)
}

func performUpdateRequest(){
	todo := Todo{
		UserId: 2345,
		Title: "Mohak Singh Go",
		Completed: false,
	}

	//convert the Todo struct to JSON
	jsonData,err := json.Marshal(todo)
	if err!=nil{
		fmt.Println("Error marshalling",err)
		return
	}

	//convert json to string
	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	const myUrl="https://jsonplaceholder.typicode.com/todos/1"

	//create put request
	req,err := http.NewRequest(http.MethodPut,myUrl,jsonReader)
	if err!=nil{
		fmt.Println("Error creating request",err)
		return
	}
	req.Header.Set("Content-Type","application/json")

	//send the request
	client := http.Client{}
	res,err := client.Do(req)
	if err !=nil{
		fmt.Println("Error sending request",err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response: ",string(data))
	fmt.Println("Status is : ",res.Status)
}

func performDeleteRequest(){
	const myUrl="https://jsonplaceholder.typicode.com/todos/1"

	//create HTTP Req
	req,err := http.NewRequest(http.MethodDelete,myUrl,nil)
	if err!=nil{
		fmt.Println("Error creating request",err)
		return
	}


	//send the request
	client := http.Client{}
	res,err := client.Do(req)
	if err!=nil{
		fmt.Println("Error sending request",err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Status is : ",res.Status)

}

func main() {
	fmt.Println("Learning CRUD operations in GoLang")
	// perfomGetRequest()
	// performPostRequest()
	// performUpdateRequest()
	performDeleteRequest()
}
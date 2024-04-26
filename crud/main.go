package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct{
	UserId int `json:"userID"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func main() {
	fmt.Println("Learning CRUD operations in GoLang")
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
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct{
	Name string	`json:"name"`
	Age int `json:"age"`
	IsAdult bool 	`json:"isAdult"`
}

func main() {
	fmt.Println("We are learnign JSON")
	person := Person{Name: "John",Age:34,IsAdult: true}
	fmt.Println("Person Data is : ",person)

	//convert person into JSON encoding(marshalling)
	jsonData,err := json.Marshal(person)
	if err!=nil{
		fmt.Println("Error in marshalling",err)
		return
	}
	fmt.Println("person Data is: ",string(jsonData))

	//decoding(Unmarshalling) JSON data into a struct
	var personData Person
	err=json.Unmarshal(jsonData,&personData)
	if err!=nil{
		fmt.Println("Error unmarshalling ",err)
		return
	}
	fmt.Println("person data is: ",person)

}


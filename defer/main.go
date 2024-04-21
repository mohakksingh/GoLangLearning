package main

import "fmt"

func add(a,b int)int{
	return a+b
}

func main() {
	fmt.Println("Starting of the program")
	data:= add(10,20)
	defer fmt.Println("Data is ",data)
	defer fmt.Println("Middle of the program")
	fmt.Println("End of the program")
}
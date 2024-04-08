package main

import "fmt"

func simpleFunction(){
	fmt.Println("This is a simple function")
}

func add(a,b int)(result int){
	result= a+b
	return result
}

func multiply(a,b int)(result int){
	result =a*b
	return
}

func main() {
	fmt.Println("We are learning functions in GoLang")
	simpleFunction()

	ans:=add(3,4)
	fmt.Println("addition is ",ans)

	product:=multiply(3,4)
	fmt.Println("product is",product)
}
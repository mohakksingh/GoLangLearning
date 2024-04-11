package main

import "fmt"

func main() {
	fmt.Println("we are learning Array in GOland")

	// var name[5]string
	// name[0] = "Raj"
	// name[2] = "Rahul"

	// fmt.Println("Names of Person is :",name)

	// var numbers = [5]int{1,2,3,4,5}
	// fmt.Println("Numbers are :",numbers)
	// fmt.Println("Length of Numbers array is :",len(numbers))
	// fmt.Println("value of name at 2 index is:",name[2])

	var price[5]string
	price[2]="Prince"
	price[0]="Aakash"
	fmt.Println("Price is :",price)
	fmt.Printf("Value of price is %q",price)
}
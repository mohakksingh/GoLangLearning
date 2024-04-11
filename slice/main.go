package main

import "fmt"

func main() {
	// numbers := []int{1, 2, 3}
	// numbers = append(numbers,45,34,5,34,3,54,35,345,34,5)
	// fmt.Println("Numbers",numbers)
	// fmt.Printf("Numbers of data type: %T",numbers)
	// fmt.Println("Length:",len(numbers))

	// name:= []string{}
	// fmt.Println("name:",name)

	numbers:= make([]int,3,5)

	numbers=append(numbers, 4,5,7)

	fmt.Println("Slice:",numbers)
	fmt.Println("Length:",len(numbers))
	fmt.Println("Capacity:",cap(numbers))

	stock := make([]int,0)
	fmt.Println("Slice:",stock)
	fmt.Println("Length:",len(stock))
	fmt.Println("Capacity:",cap(stock))



}
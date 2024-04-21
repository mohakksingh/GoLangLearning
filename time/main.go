package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current time is:", currentTime)
	fmt.Printf("Type of currentTime %T\n",currentTime)

	formatted := currentTime.Format("2006/01/02,03:04 PM")
	fmt.Println("Formatted time is:",formatted)
}
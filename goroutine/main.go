package main

import (
	"fmt"
	"time"
)

func sayHello(){
	fmt.Println("Hello")
	time.Sleep(2000*time.Millisecond)//simulating some work
	fmt.Println("sayHello function ended execution")
}

func sayHi(){
	fmt.Println("hi mohak")
	time.Sleep(1000*time.Millisecond)//simulating some work
}

func main(){
	fmt.Println("learning go routines")
	go sayHello()
	go sayHi()

	//wait for a moment to allow the goroutine to finish
	time.Sleep(1000*time.Millisecond)
}
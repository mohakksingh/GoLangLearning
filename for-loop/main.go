package main

import "fmt"

func main() {
	for i := 0; i < 1; i++ {
		fmt.Println("Numbers is",i)
	}

	counter :=0
	for{
		fmt.Println("Inifinite Loop")
		counter++
		if(counter==1){
			break
		}
	}

	numbers:=[]int{1,2,3,4,5}
	for index,value :=range numbers{
		fmt.Printf("Index of Numbers is %d, and value is %d\n ",index,value)
	}

	data:="Hello world"
	for index,char :=range data{
		fmt.Printf("Index of a data is %d, and value is %c\n",index,char)
	}
}
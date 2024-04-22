package main

import (
	"fmt"
	"os"
)

func main() {
	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error creating file",err)
	// 	return
	// }
	// defer file.Close()

	// content := "hello from go"
	// byte ,errors:=io.WriteString(file,content+"\n")
	// fmt.Println("Bytes written are: ",byte)
	// if errors != nil{
	// 	fmt.Println("Error while writing file: ",errors)
	// 	return
	// }

	// fmt.Println("File created")

	// file,err :=os.Open("example.txt")
	// if err != nil {
	// 		fmt.Println("Error creating file",err)
	// 		return
	// 	}
	// defer file.Close()

	// buffer := make([]byte,1023)

	// for {
	// 	n,err := file.Read(buffer)
	// 	if err == io.EOF{
	// 		break
	// 	}
	// 	if err!=nil{
	// 		fmt.Println("Error while reading file: ",err)
	// 		return
	// 	}

	// 	fmt.Println(string(buffer[:n]))
	// }

	content,err :=os.ReadFile("example.txt")
	if err != nil{
		fmt.Println("Error while reading the file",err)
		return
	}
	fmt.Println(string(content))

}
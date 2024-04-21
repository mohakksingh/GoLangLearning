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

	layout_str := "2006-01-02"
	datStr := "2023-11-25"
	formatted_time,_ :=time.Parse(layout_str,datStr)
	fmt.Println("Formatted time is:",formatted_time)

	//add 1 more day in currentTime
	new_date := currentTime.Add(24*time.Hour)
	fmt.Println("New date is:",new_date)
	formatted_new_date:=new_date.Format("2006/01/02 Monday")
	fmt.Println("Formatted new date is:",formatted_new_date)
}	

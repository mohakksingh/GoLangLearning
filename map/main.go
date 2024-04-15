package main

import "fmt"

func main() {

	//name <-> grage
	studentGrades := make(map[string]int)

	studentGrades["Prince"] = 100
	studentGrades["Alice"] = 95
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 95

	fmt.Println("Marks of bob:",studentGrades["Bob"])
	studentGrades["Bob"]=100
	fmt.Println("new Marks of bob:",studentGrades["Bob"])
	
	delete(studentGrades,"Bob")
	fmt.Println("Marks of bob:",studentGrades["Bob"])

	//checking if a key exists
	grades,exists := studentGrades["David"]
	fmt.Println("Grades of David",grades)
	fmt.Println("David exits",exists)

	// fmt.Println("Marks of David:",studentGrades["David"])

	//checking if a key exists
	Grades,Exists := studentGrades["Prince"]
	fmt.Println("Grades of David",Grades)
	fmt.Println("David exits",Exists)

	for index,value := range studentGrades{
		fmt.Printf("Key is %s and marks is %d\n",index,value)
	}

	person := map[string]int{
		"Alice":90,
		"Bob":85,
		"Charlie":95,
	}

	for index,value := range person{
		fmt.Printf("---------------Key is %s and marks is %d\n",index,value)
	}
}
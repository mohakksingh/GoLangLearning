package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct{
	Email string
	Phone string
	Fax string
}

type Address struct{
	House int
	Area string
	State string
}

type Employee struct{
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}


func main() {
	var prince Person
	// fmt.Println("Prince person: ",prince)
	prince.FirstName="Prince"
	prince.LastName="Agarwal"
	prince.Age=24
	// fmt.Println("Prince person: ",prince)

	Person1 := Person{
		FirstName: "Aakash",
		LastName: "Singh",
		Age:26,
	}
	fmt.Println("Person1 person: ",Person1)

	var person2=new(Person)
	person2.FirstName="Simran"
	person2.LastName="Agarwal"
	person2.Age=26

	// fmt.Println("person 2: ",person2)

	// fmt.Println("Age of prince is: ",prince.Age)

	var employee1 Employee
	employee1.Person_Details=Person{
		FirstName: "Mohak",
		LastName: "Singh",
		Age: 26,
	}
	employee1.Person_Contact.Email="contact@gmail.com"
	employee1.Person_Contact.Phone="28593485"
	employee1.Person_Contact.Fax="sfjk323jlrj324"

	employee1.Person_Address=Address{
		House:12,
		Area: "Delhi",
		State: "Delhi",
	}
	fmt.Println("employee1 : ",employee1)

}
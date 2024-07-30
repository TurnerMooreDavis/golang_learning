package main

import (
	"fmt"
	"reflect"
)

// Define the User struct
type User struct {
	id    int
	age   int
	email string
}

func main() {
	// arr1 := [3]int{1, 2, 3}
	users := &[3]User{{
		id:    1,
		age:   23,
		email: "Thing@thing.com",
	},
		{
			id:    2,
			age:   24,
			email: "Thing2@thing.com",
		},
		{
			id:    3,
			age:   4,
			email: "Thing3@thing.com",
		},
	}

	// Loop over the list of users and print their attributes

	for idx, user := range users {
		fmt.Printf("index is %d \n", idx)
		fmt.Printf("ID: %d\n", user.id)
		fmt.Printf("Name: %d\n", user.age)
		fmt.Printf("Email: %s\n", user.email)

		var typeof = reflect.TypeOf(user)
		fmt.Println(typeof) // Print type
		fmt.Println()       // Print a blank line between users
		for i := 0; i < typeof.NumField(); i++ {
			field := typeof.Field(i)
			fmt.Printf("Field Name: %s, Field Type: %s\n", field.Name, field.Type)
		}
	}
}

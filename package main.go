package main

import (
	"Assignment1/search"
	"fmt"
)

func main() {
	// Read student code from user input
	var studentCode string
	fmt.Print("Enter student code : ")
	_, err := fmt.Scan(&studentCode)
	if err != nil {
		fmt.Println("Error reading student code", err)
		return
	}

	//search the student by student code
	student, err := search.FindStudentByCode(studentCode)
	if err != nil {
		fmt.Println("Error finding student:", err)
		return
	}

	fmt.Printf("ID: %s\n", student.ID)
	fmt.Printf("Code: %s\n", student.Code)
	fmt.Printf("Name: %s\n", student.Name)
	fmt.Printf("Address: %s\n", student.Address)
	fmt.Printf("Occupation: %s\n", student.Occupation)
	fmt.Printf("Reason: %s\n", student.Reason)
}

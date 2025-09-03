package main

import (
	"fmt"
)

// var email string = "prungprechskun_j@su.ac.th"
func main() {
	// var name string = "Jarudet"
	var age int = 20

	email := "prungprechskun_j@su.ac.th"
	gpa := 2.50

	firstName, lastname := "Jarudet", "Prungprechskun"

	fmt.Printf("Name %s %s, age %d, email %s, gpa %.2f\n", firstName, lastname, age, email, gpa)

}

package main

import (
	"fmt"
)

func main(){
	// The array
	var nameArray [2]string

	// Values
	nameArray[0] = "John"
	nameArray[1] = "Alex"
	// Print the array
	fmt.Println(nameArray)
	fmt.Println("\nThe first name is",nameArray[0] )
	fmt.Println("\nThe second name is",nameArray[1] )
	fmt.Println(len(nameArray))
}

// Main Package
// Entry point to the program
package main 

import "go/generics"

func main() {
	intArr := []int{1, 2, 3, 4, 5}
	strArr := []string{"apple", "ball", "cat", "dog", "elephant"}

	generics.PrintArray(intArr)
	generics.PrintArray(strArr)
}
// Main Package
// Entry point to the program
package main 

import (
	"fmt"
	"go/generics"
)

func main() {
	// Integer Stack
	intStack := generics.Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	fmt.Println("Stack Size: ", intStack.Size())
	fmt.Println("Is Stack Empty: ", intStack.IsEmpty())
	val, isPopped := intStack.Pop()
	if isPopped == true {
		fmt.Println("Popped: ", val)
	} else {
		fmt.Println("Stack is Empty")
	}
	intStack.PrintStack()

	// String Stack
	stringStack := generics.Stack[string]{}
	stringStack.Push("Hello World")
	stringStack.Push("Good Morning")
	stringStack.Push("Hello Everyone")
	stringStack.PrintStack()

	// Float Stack
	floatStack := generics.Stack[float32]{}

	fmt.Printf("intStack Data type: %T\n", intStack.StackType())
	fmt.Printf("stringStack Data type: %T\n", stringStack.StackType())
	fmt.Printf("floatStack Data type: %T\n", floatStack.StackType())
}
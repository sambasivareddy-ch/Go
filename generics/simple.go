package generics

import "fmt"

/*
	[T any] => defines a generic type which accepts "any" datatype.
				Basically T can be of any data type
	Params:
		arr []T => defines the array of any data 
*/
func PrintArray[T any] (arr []T) {
	for _, value := range arr {
		fmt.Println(value)
	}
}
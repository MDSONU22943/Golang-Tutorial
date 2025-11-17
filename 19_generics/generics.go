package main

import "fmt"

// func printSLice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }


// func printSLice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printSLice[T any](items []T) {
// 		for _, item := range items {
// 			fmt.Println(item)
// 		}
// 	}

type stack struct {
	elements []int
}

func main() {

	myStack := stack{
		elements: []int{1,2,3},
	}
	fmt.Println(myStack)
	// nums := []int{1,2,3,4}
	// printSLice(nums)

	// names := []string{"golang", "java"}
	// printSLice(names)

	// printSLice([]int{1,2,3})




}
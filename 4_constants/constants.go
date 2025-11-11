package main

import "fmt"

// const age=98
func main(){
	const name string ="golang"

	// name="js"
	// const age=30

	// age=87

	// fmt.Println(age)

	const (
		port = 5000
		host = "localhost"
	)
	fmt.Print(port,host)
}
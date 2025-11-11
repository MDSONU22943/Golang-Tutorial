package main

import (
	"fmt"
	
)

// maps --> hash, object, dict
func main() {
	// creating map

	// m := make(map[string]string)

	// setting an element
	// m["name"]="golang"
	// m["area"]="backend"


	// get an element
	// fmt.Println(m["name"], m["area"])

	//IMP: if key does not exists in the map the it returns zero value
	// fmt.Println(m["phone"])

	// m := make(map[string]int)
	// m["age"]=30
	// m["price"]=50
	// fmt.Println(m["phone"])

	// fmt.Println(len(m))

	// delete(m, "price")
	// clear(m)
	// fmt.Println(m)

	// fmt.Println(m)
	
	m := map[string]int{"price":40, "phone":3}

	k, ok := m["price"]

	fmt.Println(k)

	if ok {
		fmt.Println("all ok")
	}else {
		fmt.Println("not ok")
	}





}

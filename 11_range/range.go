package main

import "fmt"

// iterating over data structures
func main(){

	// nums := []int{6,7,8}

	// for i:=0 i<len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }
	// sum := 0
	// for i, num := range nums {
	// 	sum = sum+num
	// 	fmt.Println(num, i)
	// }
	// fmt.Println(sum)


	// m := map[string]string{"fname":"john", "lname":"doe"}

	// for k, v := range m {
	// 	fmt.Println(k,v)
	// }

	// unicode code
	for i, c := range "golang" {
		fmt.Println(i,string(c))
	}
}
package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 2, 3} // slice are like dynamic array (memory)
	age := nums

	age[0] = 24

	fmt.Println(nums)
	fmt.Println(age)
}

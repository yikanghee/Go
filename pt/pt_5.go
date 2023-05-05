package main

import "fmt"

func main() {
	var numbers1 []int
	numbers1 = append(numbers1, 1)

	numbers2 := []int{1, 2, 3}
	fmt.Println(numbers1)
	fmt.Println(numbers2)
}

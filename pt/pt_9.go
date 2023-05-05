package main

import "fmt"

func main() {
	myMap := map[int]string{}
	myMap[1] = "Seoul"
	myMap[2] = "London"

	for key, value := range myMap {
		fmt.Println(key, value)
	}
}

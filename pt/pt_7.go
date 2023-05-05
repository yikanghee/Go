package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(input, "")
	for _, value := range inputs {
		fmt.Println(value)
	}
}

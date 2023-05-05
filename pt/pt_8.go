package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}

	var numbers []int
	strNumbers := strings.Split(input, ",")
	for _, strNum := range strNumbers {
		strNum, _ := strconv.Atoi(strNum)
		numbers = append(numbers, strNum)
	}
	fmt.Println(numbers[2:5])

}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getGrade2(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func main() {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}

	strNums := strings.Split(input, ",")

	for _, str := range strNums {
		num, _ := strconv.Atoi(str)
		fmt.Println(getGrade2(num))
	}
}

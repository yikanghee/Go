package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getGrade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
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
		fmt.Println(getGrade(num))
	}

}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Member struct {
	Age    int
	Salary int
}

func (member *Member) NextYearAge() {
	member.Age++
}

func (member Member) BonusSalary() {
	member.Salary += 10000
	fmt.Println(fmt.Sprintf("보너스가 있을 때, salary : %d", member.Salary))
}

func main() {
	age, salary := getInput()
	member := &Member{age, salary}

	fmt.Println(fmt.Sprintf("2023년 5월일 때, 나이 : %d", member.Age))
	member.NextYearAge()
	fmt.Println(fmt.Sprintf("20224년 5월일 때, 나이 : %d", member.Age))

	member.BonusSalary()
	fmt.Println(fmt.Sprintf("일반적인 salary : %d", member.Salary))
}

func getInput() (int, int) {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(input, ",")
	age, err := strconv.Atoi(inputs[0])
	if err != nil {
		panic(err)
	}
	salary, err := strconv.Atoi(inputs[1])
	if err != nil {
		panic(nil)
	}
	return age, salary
}

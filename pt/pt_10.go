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
	/*
		myMap -> map[string]struct{}{}
		keys -> strings.Split(input, "")
		myMap[key] <- keys = struct{}{}
	*/
	myMap := make(map[string]struct{})
	keys := strings.Split(input, "")
	for _, key := range keys {
		myMap[key] = struct{}{}
	}

	/*
		expectedKeys -> []string{"", "", ""}
		exitedKeys [] string
		myMap[key] <- struct{}{}
		_, ok := myMap[key] <- expectedKeys => true, false
		if ok {
			exitedKeys = append(exitedKeys, key)
		}
	*/
	expectedKeys := []string{"a", "b", "c", "d", "e"}
	var exitedKeys []string
	for _, key := range expectedKeys {
		_, ok := myMap[key]
		fmt.Println(ok)
		if ok {
			exitedKeys = append(exitedKeys, key)
		}
	}
	fmt.Println(exitedKeys)

}

package main

import (
	"fmt"
	"sort"
	"strings"
)

func reorganizeString(s string) (string, error) {
	if len(s) < 1 || len(s) > 500 {
		return "", fmt.Errorf("Error: Invalid input length. Please provide a string with length between 1 and 500.")
	}

	if strings.ToLower(s) != s {
		return "", fmt.Errorf("Error: Uppercase characters are not allowed. Please provide a lowercase string.")
	}

	counter := make(map[rune]int)
	for _, char := range s {
		counter[char]++
	}

	chars := make([][2]interface{}, 0)
	for char, count := range counter {
		chars = append(chars, [2]interface{}{char, count})
	}
	sort.Slice(chars, func(i, j int) bool {
		return chars[i][1].(int) > chars[j][1].(int)
	})

	result := make([]rune, len(s))
	index := 0
	for _, charCount := range chars {
		char := charCount[0].(rune)
		count := charCount[1].(int)
		if count > (len(s)+1)/2 {
			return "", nil
		}
		for count > 0 {
			result[index] = char
			index += 2
			if index >= len(s) {
				index = 1
			}
			count--
		}
	}
	return string(result), nil
}

func main() {
	var userInput string
	fmt.Print("Enter a lowercase string (1 to 500 characters): ")
	fmt.Scan(&userInput)

	result, err := reorganizeString(userInput)
	if err != nil {
		fmt.Println(err)
	} else if result == "" {
		fmt.Println("Result:=\"\"")
	} else {
		fmt.Println("Result:", "\""+result+"\"")
	}
}

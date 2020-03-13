package main

import "fmt"

var ans = []string{}

func getSubsequent(str string, output string, answer *[]string) []string {
	if len(str) == 0 {
		*answer = append(*answer, output)
		return *answer
	}
	getSubsequent(str[1:], output+string(str[0]), answer)
	getSubsequent(str[1:], output, answer)
	return *answer
}

func main() {
	str := "abcde"
	output := ""
	fmt.Println(getSubsequent(str, output, &ans))
}

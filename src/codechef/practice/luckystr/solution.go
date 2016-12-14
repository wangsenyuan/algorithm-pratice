package main

import (
	"fmt"
	"strings"
)

func main() {
	k, n := 0, 0
	fmt.Scanf("%d %d\n", &k, &n)

	as := make([]string, k)
	for i := 0; i < k; i++ {
		as[i] = readLine()
	}

	for i := 0; i < n; i++ {
		str := readLine()
		if checkGoodStr(string(str), as) {
			fmt.Println("Good")
		} else {
			fmt.Println("Bad")
		}
	}
}

func readLine() string {
	var str string
	fmt.Scanf("%s\n", &str)
	return str
}

func checkGoodStr(str string, as []string) bool {
	if len(str) >= 47 {
		return true
	}

	for _, a := range as {
		if strings.Contains(str, a) {
			return true
		}
	}
	return false
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "172162541"
	rs := restoreIpAddresses(str)

	for _, r := range rs {
		fmt.Println(r)
	}
}

func restoreIpAddresses(s string) []string {
	n := len(s)
	var result []string

	var process func(i int, j int, path []string)

	process = func(i int, j int, path []string) {
		if i-j > 3 || len(path) > 4 {
			return
		}
		if j == len(s) {
			if len(path) == 4 {
				result = append(result, strings.Join(path, "."))
			}
			return
		}

		process(i+1, j, path)

		if j <= n-1 && i-j == 1 ||
			j <= n-2 && i-j == 2 && s[j] != '0' ||
			j <= n-3 && s[j] == '1' ||
			j <= n-3 && s[j] == '2' && s[j+1] == '5' && s[j+2] <= '5' ||
			j <= n-3 && s[j] == '2' && s[j+1] < '5' {
			process(i+1, i, append(path, s[j:i]))
		}

	}

	process(1, 0, make([]string, 0, 5))

	return result
}

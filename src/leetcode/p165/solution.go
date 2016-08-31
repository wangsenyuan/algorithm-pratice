package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(compareVersion("1.2.3.4.5", "1.2.3.4.5.0.0"))
}

func compareVersion(version1 string, version2 string) int {
	if len(version1) == 0 && len(version2) == 0 {
		return 0
	}
	a, i := number(version1)
	b, j := number(version2)

	if a > b {
		return 1
	}
	if a < b {
		return -1
	}

	return compareVersion(version1[i:], version2[j:])
}

func number(s string) (int, int) {
	if len(s) == 0 {
		return 0, 0
	}

	i := strings.Index(s, ".")

	if i < 0 {
		x, _ := strconv.Atoi(s)
		return x, len(s)
	}

	x, _ := strconv.Atoi(s[:i])
	return x, min(len(s), i+1)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
	a := readString(reader)
	b := readString(reader)
	res := solve(a, b)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func solve(a string, b string) int {

	// x = i-2位置处的最优解
	// y = i-1相同时的最优解
	x, y := 0, 0
	if a[0] != b[0] {
		// flip a[0]
		y++
	}
	n := len(a)
	for i := 1; i < n; i++ {
		if a[i] == b[i] {
			x = y
			// y keeps
			continue
		}
		if a[i] == b[i-1] && a[i-1] == b[i] {
			// swap
			x, y = y, x+1
			continue
		}
		// flip
		x, y = y, y+1
	}

	return y
}

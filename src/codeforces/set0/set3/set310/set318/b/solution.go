package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)
	res := solve(s)

	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

const heavy = "heavy"
const metal = "metal"

func solve(s string) (res int) {
	var cnt int

	for i := 0; i+5 <= len(s); i++ {
		if s[i:i+5] == heavy {
			cnt++
		}
		if s[i:i+5] == metal {
			res += cnt
		}
	}
	return res
}

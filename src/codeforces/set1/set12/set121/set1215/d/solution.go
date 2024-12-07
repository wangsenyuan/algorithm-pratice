package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
	s := readString(reader)
	res := solve(s)
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

const Monocarp = "Monocarp"
const Bicarp = "Bicarp"

func solve(s string) string {
	var sum int
	cnt := []int{0, 0}

	n := len(s) / 2
	for i := 0; i < n; i++ {
		if s[i] == '?' {
			cnt[0]++
		} else {
			x := int(s[i] - '0')
			sum += x
		}
	}

	for i := n; i < 2*n; i++ {
		if s[i] == '?' {
			cnt[1]++
		} else {
			x := int(s[i] - '0')
			sum -= x
		}
	}

	lo := sum - cnt[1]*9
	hi := sum + cnt[0]*9

	if lo+hi == 0 {
		return Bicarp
	}

	return Monocarp
}

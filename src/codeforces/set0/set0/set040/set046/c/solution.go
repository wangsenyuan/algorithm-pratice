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
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func solve(s string) int {
	n := len(s)

	var cnt int

	for _, x := range []byte(s) {
		if x == 'H' {
			cnt++
		}
	}
	// 也就是要找出一段连续的, 且长度为cnt区间， 且这个区间内,T的数量最少
	ans := n

	var cnt2 int
	for i := 0; i < cnt+n; i++ {

		if s[i%n] == 'T' {
			cnt2++
		}
		if i >= cnt {
			if s[(i-cnt)%n] == 'T' {
				cnt2--
			}
		}
		if i >= cnt-1 {
			ans = min(ans, cnt2)
		}
	}
	return ans
}

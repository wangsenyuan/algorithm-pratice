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
	if res < 0 {
		fmt.Println("Impossible")
		return
	}
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

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(s string) int {
	n := len(s)

	if n <= 3 || check(s) {
		return -1
	}
	if check1(s) {
		return 1
	}

	return 2
}

// 检查是否是同一个字符
func check(s string) bool {
	n := len(s)
	if s[0] == s[1] {
		var cnt int
		for i := 0; i < n; i++ {
			if s[i] == s[0] {
				cnt++
			}
		}
		return cnt >= n-1
	}
	// s[0] != s[1]
	cnt := make([]int, 2)
	for i := 0; i < n; i++ {
		if s[i] == s[0] {
			cnt[0]++
		} else if s[i] == s[1] {
			cnt[1]++
		}
	}
	return cnt[0] >= n-1 || cnt[1] >= n-1
}

func check1(s string) bool {
	n := len(s)
	s += s
	for i := 1; i < n; i++ {
		l, r := i, i+n-1
		for l < r && s[l] == s[r] {
			l++
			r--
		}
		if l >= r && s[i:i+n] != s[0:n] {
			return true
		}
	}
	return false
}

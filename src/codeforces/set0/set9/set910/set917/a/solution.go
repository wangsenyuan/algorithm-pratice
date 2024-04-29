package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)

	fmt.Println(solve(s))
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
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

	fp := make([][]int32, n)

	for i := 0; i < n; i++ {
		fp[i] = make([]int32, n)
	}

	for i := 0; i < n; i++ {
		var level int
		for j := i; j < n; j++ {
			if s[j] == '(' || s[j] == '?' {
				level++
			} else {
				level--
			}

			if level < 0 {
				break
			}
			if (j-i+1)%2 == 0 {
				fp[i][j]++
			}
		}
		level = 0
		for j := i; j >= 0; j-- {
			if s[j] == ')' || s[j] == '?' {
				level++
			} else {
				level--
			}
			if level < 0 {
				break
			}
			if (i-j+1)%2 == 0 {
				fp[j][i]++
			}
		}
	}

	var res int

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if fp[i][j] == 2 {
				res++
			}
		}
	}

	return res
}

func check(level int, marks int) bool {
	return marks >= abs(level)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

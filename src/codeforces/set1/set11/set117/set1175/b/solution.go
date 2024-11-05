package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	statements := make([]string, n)
	for i := 0; i < n; i++ {
		statements[i] = readString(reader)
	}
	res := solve(statements)
	if res < 0 {
		fmt.Println("OVERFLOW!!!")
	} else {
		fmt.Println(res)
	}
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

const inf = 1<<32 - 1

func solve(statements []string) int {
	// 多个add可以合并成一个 add cnt
	n := len(statements)
	stack := make([]int, n)
	var top int

	var ans int

	for _, cur := range statements {
		if cur == "add" {
			if top > 0 {
				if stack[top-1] > inf || stack[top-1] < 0 {
					return -1
				}
				ans += stack[top-1]
			} else {
				ans++
			}
		} else if cur == "end" {
			top--
		} else {
			// for
			var v int
			readInt([]byte(cur), 4, &v)
			if top > 0 {
				v *= stack[top-1]
			}
			if v < 0 || v > inf {
				v = -1
			}
			stack[top] = v
			top++
		}
		if ans > inf {
			return -1
		}
	}

	return ans
}

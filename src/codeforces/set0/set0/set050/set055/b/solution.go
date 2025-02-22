package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nums := readNNums(reader, 4)
	s := readString(reader)
	res := solve(nums, s)
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

const inf = 1 << 60

func solve(nums []int, s string) int {
	ops := strings.Split(s, " ")

	calc := func(a int, b int, c string) int {
		if c == "+" {
			return a + b
		}
		return a * b
	}

	var dfs func(arr []int) int

	dfs = func(arr []int) int {
		if len(arr) == 1 {
			return arr[0]
		}
		j := 4 - len(arr)
		// 选择两个数出来
		res := inf
		for u := 0; u < len(arr); u++ {
			for v := u + 1; v < len(arr); v++ {
				w := calc(arr[u], arr[v], ops[j])
				tmp := []int{w}
				for i := 0; i < len(arr); i++ {
					if i != u && i != v {
						tmp = append(tmp, arr[i])
					}
				}
				res = min(res, dfs(tmp))
			}
		}
		return res
	}

	return dfs(nums)
}

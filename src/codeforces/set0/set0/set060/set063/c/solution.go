package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func process(reader *bufio.Reader) string {
	n := readNum(reader)
	guesses := make([]string, n)
	x := make([]int, n)
	y := make([]int, n)
	for i := range n {
		s := readString(reader)
		ss := strings.Split(s, " ")
		guesses[i] = ss[0]
		x[i], _ = strconv.Atoi(ss[1])
		y[i], _ = strconv.Atoi(ss[2])
	}

	return solve(guesses, x, y)
}

func solve(guess []string, x []int, y []int) string {
	n := len(guess)

	pos := make([]int, 10)

	play := func(correct string, cur string) (x int, y int) {
		clear(pos)

		for i := 0; i < 4; i++ {
			pos[int(correct[i]-'0')] = i + 1
		}

		for i := 0; i < 4; i++ {
			j := pos[int(cur[i]-'0')]
			if i+1 == j {
				x++
			} else if j > 0 {
				y++
			}
		}
		return
	}

	check := func(num string) bool {
		// 和现有的是否匹配
		for i := range n {
			u, v := play(num, guess[i])
			if u != x[i] || v != y[i] {
				return false
			}
		}
		return true
	}
	buf := make([]byte, 4)
	var res []string
	var dfs func(i int, flag int)
	dfs = func(i int, flag int) {
		if i == 4 {
			if len(res) < 2 && check(string(buf)) {
				res = append(res, string(buf))
			}
			return
		}
		for v := 0; v < 10; v++ {
			if (flag>>v)&1 == 0 {
				buf[i] = byte(v + '0')
				dfs(i+1, flag|(1<<v))
				if len(res) >= 2 {
					// no need to process any more
					return
				}
			}
		}
	}

	dfs(0, 0)

	if len(res) == 0 {
		return "Incorrect data"
	}
	if len(res) > 1 {
		return "Need more data"
	}

	return res[0]
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = readString(reader)
	}
	res := solve(words)
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

func solve(words []string) int {
	//n := len(words)

	// 一个要考虑它最终的平衡性，还要考虑它的最小的level
	// 如果一个字符既有未匹配的右括号，又有未匹配的左括号，那么这个是无效的
	var zeros int
	l := make(map[int]int)
	r := make(map[int]int)

	for _, w := range words {
		var u, v int
		for i := 0; i < len(w); i++ {
			if w[i] == '(' {
				u++
			} else {
				if u == 0 {
					v++
				} else {
					u--
				}
			}
		}
		if u > 0 && v > 0 {
			// invalid
			continue
		}

		if u == 0 && v == 0 {
			zeros++
			continue
		}

		if v > 0 {
			r[v]++
		} else {
			l[u]++
		}
	}

	ans := zeros / 2
	for k, v := range l {
		w := r[k]
		ans += min(v, w)
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

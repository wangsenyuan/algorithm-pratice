package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)
	n := readNum(reader)
	queries := make([][]int, n)
	for i := 0; i < n; i++ {
		queries[i] = readNNums(reader, 2)
	}

	res := solve(s, queries)

	var buf bytes.Buffer

	for _, x := range res {
		if x {
			buf.WriteString("Yes\n")
		} else {
			buf.WriteString("No\n")
		}
	}
	fmt.Print(buf.String())
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

func solve(s string, queries [][]int) []bool {
	n := len(s)
	cnt := make([][]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = make([]int, 26)
		if i > 0 {
			copy(cnt[i], cnt[i-1])
		}
		cnt[i][int(s[i]-'a')]++
	}

	get := func(l int, r int, x int) int {
		res := cnt[r][x]
		if l > 0 {
			res -= cnt[l-1][x]
		}
		return res
	}

	check := func(l int, r int) bool {
		if l == r {
			return true
		}
		if s[l] != s[r] {
			return true
		}
		var arr []int
		for i := 0; i < 26; i++ {
			tmp := get(l, r, i)
			if tmp > 0 {
				arr = append(arr, tmp)
			}
			if len(arr) >= 3 {
				return true
			}
		}

		return false
	}

	ans := make([]bool, len(queries))

	for i, cur := range queries {
		l, r := cur[0], cur[1]
		l--
		r--
		ans[i] = check(l, r)
	}
	return ans
}

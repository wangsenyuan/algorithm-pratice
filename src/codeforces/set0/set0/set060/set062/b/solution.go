package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
		if i > 0 {
			buf.WriteByte('\n')
		}
		buf.WriteString(strconv.Itoa(res[i]))
	}
	fmt.Println(buf.String())
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

func process(reader *bufio.Reader) []int {
	n, _ := readTwoNums(reader)
	s := readString(reader)
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = readString(reader)
	}
	return solve(s, words)
}

func solve(s string, words []string) []int {
	k := len(s)
	n := len(words)

	pre := make([][]int, k)
	for i := range k {
		pre[i] = make([]int, 26)
		if i > 0 {
			copy(pre[i], pre[i-1])
		} else {
			for j := range 26 {
				pre[i][j] = -1
			}
		}
		pre[i][s[i]-'a'] = i
	}

	suf := make([][]int, k)
	for i := k - 1; i >= 0; i-- {
		suf[i] = make([]int, 26)
		if i < k-1 {
			copy(suf[i], suf[i+1])
		} else {
			for j := range 26 {
				suf[i][j] = k
			}
		}
		suf[i][s[i]-'a'] = i
	}

	check := func(c string) int {
		var res int
		var i int
		for i < len(c) && i < k {
			x := int(c[i] - 'a')
			if pre[i][x] >= 0 && suf[i][x] < k {
				res += min(i-pre[i][x], suf[i][x]-i)
			} else if pre[i][x] >= 0 {
				res += i - pre[i][x]
			} else if suf[i][x] < k {
				res += suf[i][x] - i
			} else {
				res += len(c)
			}
			i++
		}
		for i < len(c) {
			x := int(c[i] - 'a')
			if pre[k-1][x] >= 0 {
				res += i - pre[k-1][x]
			} else {
				res += len(c)
			}
			i++
		}
		return res
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		word := words[i]
		ans[i] = check(word)
	}
	return ans
}

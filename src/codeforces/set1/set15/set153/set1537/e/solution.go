package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	s := readString(reader)[:n]
	res := solve(s, k)
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

func solve(s string, k int) string {

	z := zFunction(s)
	n := len(s)

	cur := 1

	get := func(m int) string {
		buf := make([]byte, k)
		for i := 0; i < k; i++ {
			buf[i] = s[i%m]
		}
		return string(buf)
	}

	for i := 1; i < n; i++ {
		if s[i] > s[i%cur] {
			return get(cur)
		}
		if s[i] < s[i%cur] {
			cur = i + 1
			continue
		}
		off := i - cur + 1
		if off == cur {
			cur = i + 1
			continue
		}
		if z[off] < cur-off {
			if cur+off+z[off] >= k {
				cur = i + 1
				continue
			}
			if s[off+z[off]] > s[z[off]] {
				cur = i + 1
			}
			continue
		}
		if z[cur-off] < off {
			if 2*cur+z[cur-off] >= k {
				cur = i + 1
				continue
			}
			if s[cur-off+z[cur-off]] < s[z[cur-off]] {
				cur = i + 1
				continue
			}
		}
		cur = i + 1
	}

	return get(cur)
}

func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	return z
}

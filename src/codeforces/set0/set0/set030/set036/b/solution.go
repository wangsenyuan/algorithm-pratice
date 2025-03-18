package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	r, _ := os.Open("input.txt")
	defer r.Close()
	w, _ := os.Create("output.txt")
	defer w.Close()
	reader := bufio.NewReader(r)
	res := process(reader)
	var buf bytes.Buffer

	for _, cur := range res {
		buf.WriteString(cur)
		buf.WriteByte('\n')
	}
	fmt.Fprint(w, buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	if len(s) == 0 {
		return s
	}
	if s[len(s)-1] == '\n' || s[len(s)-1] == '\r' {
		return s[:len(s)-1]
	}
	return s
}

func process(reader *bufio.Reader) []string {
	n, k := readTwoNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(n, k, a)
}

func solve(n int, k int, a []string) []string {
	// 应该计算 (i, j)对应到原来的格子的位置
	w := 1
	for i := 1; i <= k; i++ {
		w *= n
	}
	var get func(x int, y int, i int) byte
	get = func(x int, y int, i int) byte {
		if i == 1 {
			return a[x][y]
		}

		if a[x%n][y%n] == '*' {
			return '*'
		}

		return get(x/n, y/n, i-1)
	}

	// w = pow(n, k)
	res := make([][]byte, w)
	for i := range w {
		res[i] = make([]byte, w)
	}
	for i := 0; i < w; i += n {
		for j := 0; j < w; j += n {
			for u := i; u < i+n; u++ {
				for v := j; v < j+n; v++ {
					res[u][v] = get(u, v, k)
				}
			}
		}
	}
	ans := make([]string, w)
	for i := range w {
		ans[i] = string(res[i])
	}
	return ans
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')
	buf.WriteTo(os.Stdout)
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

func readString(reader *bufio.Reader) (string, error) {
	s, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i], nil
		}
	}
	return s, nil
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
	n := readNum(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	return solve(a, b)
}

func solve(a []int, b []int) []int {
	n := len(a)
	s1 := make([]int, 2*n+1)
	s2 := make([]int, 2*n+1)

	ga := func(l int, r int) int {
		return s1[r] - s1[l]
	}
	gb := func(l int, r int) int {
		return s2[r] - s2[l]
	}

	check := func(a []int, b []int) []bool {
		for i := 0; i < 2*n; i++ {
			s1[i+1] = s1[i] + a[i%n]
			s2[i+1] = s2[i] + b[i%n]
		}

		ok := make([]bool, n)
		// next[i] 表示从i出发最远能到达的位置+1
		next := make([]int, 2*n+1)
		next[2*n] = 2 * n
		for i := 2*n - 1; i >= 0; i-- {
			// 默认到下一个位置
			next[i] = i + 1
			for next[i] < 2*n && ga(i, next[i]) >= gb(i, next[i]) {
				next[i] = next[next[i]]
			}
			if next[i]-i >= n {
				ok[i%n] = true
			}
		}
		return ok
	}

	ok1 := check(a, b)
	reverse(a)
	reverse(b)
	first := b[0]
	for i := 0; i < n-1; i++ {
		b[i] = b[i+1]
	}
	b[n-1] = first
	ok2 := check(a, b)
	reverse(ok2)

	var res []int

	for i := range n {
		if ok1[i] || ok2[i] {
			res = append(res, i+1)
		}
	}

	return res
}

func reverse[T any](arr []T) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

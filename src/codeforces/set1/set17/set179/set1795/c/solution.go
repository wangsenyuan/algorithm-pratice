package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		b := readNNums(reader, n)
		res := solve(a, b)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(a, b []int) []int {
	n := len(a)

	diff := make([]int, 1+n)

	// 对于a[i] 来说，要找到最大的j sum[j-i] of b > a[i]
	res := make([]int, n)

	sum := make([]int, n)
	for i := 0; i < n; i++ {
		sum[i] = b[i]
		if i > 0 {
			sum[i] += sum[i-1]
		}
	}

	for i := 0; i < n; i++ {
		var prev int
		if i > 0 {
			prev = sum[i-1]
		}
		j := search(n, func(j int) bool {
			return sum[j]-prev > a[i]
		})
		// sum[j] - prev > a[i]
		// sum[j-1] - prev <= a[i]
		if i < j {
			diff[i]++
			diff[j]--
			tmp := a[i] - (sum[j-1] - prev)
			if j < n {
				res[j] += tmp
			}
		} else {
			// b[i] >= a[i]
			res[i] += a[i]
		}
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			diff[i] += diff[i-1]
		}
		res[i] += diff[i] * b[i]
	}

	return res
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

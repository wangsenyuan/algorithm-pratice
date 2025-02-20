package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) int {
	s := readString(reader)
	t := readString(reader)
	return solve(s, t)
}

func solve(s string, t string) int {
	if s == t {
		return 0
	}
	encode := func(x string) []int {
		n := len(x)
		var arr []int
		for i := 0; i < n; {
			j := i
			for i < n && x[j] == x[i] {
				i++
			}
			arr = append(arr, i-j)
		}
		return arr
	}

	a := encode(s)
	b := encode(t)

	// s != t
	if len(a) == 1 {
		// 这个无法交换
		return -1
	}

	if s[0] != t[0] {
		a = append([]int{0}, a...)
	}
	if s[len(s)-1] != t[len(t)-1] {
		a = append(a, 0)
	}

	if len(a) < len(b) || (len(a)-len(b))%2 != 0 {
		return -1
	}

	// len(a) >= len(b)
	for i, j := 0, 0; i < len(b); i++ {
		if j == len(a) || a[j] > b[i] {
			return -1
		}
		// 这个时候，必须确定 a[j] 表示的是什么区域
		if a[j] == b[i] {
			j++
			continue
		}
		// a[j] < b[i]
		diff := b[i] - a[j]
		sum := []int{0, 0}
		k := j + 1
		for ; k < len(a); k++ {
			if (k-j)%2 == 1 {
				sum[1] += a[k]
			} else {
				sum[0] += a[k]
			}
			if sum[0] >= diff {
				break
			}
		}
		j = k + 1

		if sum[0] > diff || j >= len(a) {
			return -1
		}
		a[j] += sum[1]
	}

	return (len(a) - len(b)) / 2
}

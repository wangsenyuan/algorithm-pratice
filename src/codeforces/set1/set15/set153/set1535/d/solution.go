package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	k := readNum(reader)
	s := readString(reader)
	m := readNum(reader)
	queries := make([]string, m)
	for i := 0; i < m; i++ {
		queries[i] = readString(reader)
	}
	res := solve(s, k, queries)
	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(s string, k int, queries []string) []int {
	n := 1 << k
	// 倒过来比较好处理
	tr := make([]int, 2*n)
	for i := n; i < 2*n; i++ {
		tr[i] = 1
	}
	x := []byte(s)
	reverse(x)
	for i := n - 1; i > 0; i-- {
		if x[i-1] == '0' {
			tr[i] = tr[i*2+1]
		} else if x[i-1] == '1' {
			tr[i] = tr[i*2]
		} else {
			tr[i] = tr[i*2] + tr[i*2+1]
		}
	}

	process := func(q string) int {
		var i int
		pos := readInt([]byte(q), 0, &i)
		i = (n - 1 - i)
		c := q[pos+1]
		x[i] = c
		i++

		for i > 0 {
			if x[i-1] == '0' {
				tr[i] = tr[2*i+1]
			} else if x[i-1] == '1' {
				tr[i] = tr[2*i]
			} else {
				tr[i] = tr[2*i] + tr[2*i+1]
			}
			i >>= 1
		}

		return tr[1]
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		ans[i] = process(cur)
	}

	return ans
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

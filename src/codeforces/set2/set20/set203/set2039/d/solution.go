package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			for _, x := range res {
				buf.WriteString(fmt.Sprintf("%d ", x))
			}
			buf.WriteByte('\n')
		}
	}
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
	n, m := readTwoNums(reader)
	a := readNNums(reader, m)
	return solve(n, a)
}

func solve(n int, a []int) []int {
	sort.Ints(a)

	res := make([]int, n+1)
	m := len(a)
	for i := 1; i <= n; i++ {
		res[i] = a[m-1] + 1
	}

	for i := 1; i <= n; i++ {
		j := sort.SearchInts(a, res[i])
		if j == 0 {
			return nil
		}
		// a[j-1] < res[i]
		res[i] = a[j-1]
		for j := 2 * i; j <= n; j += i {
			res[j] = min(res[j], res[i])
		}
	}

	return res[1:]
}

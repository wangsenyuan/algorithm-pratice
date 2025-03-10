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
	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		res := process(reader)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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
	a := readNNums(reader, n)
	b := readNNums(reader, m)
	return solve(a, b)
}

func solve(a []int, b []int) []int {

	var arr []int
	for i := 1; i < len(a); i++ {
		if a[i] > a[0] {
			arr = append(arr, a[i])
		}
	}

	sort.Ints(arr)
	reverse(arr)

	n := len(b)
	sort.Ints(b)

	c := make([]int, n)

	for i, j := n-1, 0; i >= 0 && b[i] > a[0]; i-- {
		for j < len(arr) && b[i] <= arr[j] {
			j++
		}
		c[i] = j
	}

	sort.Ints(c)

	res := make([]int, n+1)

	for k := 1; k <= n; k++ {
		for x := k; x <= n; x += k {
			res[k] += c[x-1] + 1
		}
	}

	return res[1:]
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

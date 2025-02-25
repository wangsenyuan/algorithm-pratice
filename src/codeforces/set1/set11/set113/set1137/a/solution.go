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
	res := process(reader)
	for _, row := range res {
		for _, x := range row {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
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

func process(reader *bufio.Reader) [][]int {
	n, m := readTwoNums(reader)
	a := make([][]int, n)
	for i := range n {
		a[i] = readNNums(reader, m)
	}
	return solve(a)
}

func solve(a [][]int) [][]int {
	n := len(a)
	m := len(a[0])
	cols := make([]SortedUnqiueNums, m)
	arr := make([]int, n)
	for j := 0; j < m; j++ {
		for i := range n {
			arr[i] = a[i][j]
		}
		cols[j] = NewSortedUniqueNums(arr)
	}

	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, m)
		row := NewSortedUniqueNums(a[i])
		for j := 0; j < m; j++ {
			u := row.Search(a[i][j])
			v := cols[j].Search(a[i][j])
			ans[i][j] = max(len(row), len(cols[j]), u+len(cols[j])-v, v+len(row)-u)
		}
	}

	return ans
}

type SortedUnqiueNums []int

func NewSortedUniqueNums(arr []int) SortedUnqiueNums {
	n := len(arr)
	res := make([]int, n)
	copy(res, arr)
	sort.Ints(res)
	n = 0
	for i := 1; i <= len(res); i++ {
		if i == len(res) || res[i] > res[i-1] {
			res[n] = res[i-1]
			n++
		}
	}
	return SortedUnqiueNums(res[:n])
}

func (nums SortedUnqiueNums) Search(v int) int {
	return sort.SearchInts(nums, v)
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

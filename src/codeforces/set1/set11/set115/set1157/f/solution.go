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
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
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

func solve(a []int) []int {
	sort.Ints(a)

	type pair struct {
		first  int
		second int
	}
	n := len(a)

	var arr []pair
	for i := 0; i < n; {
		j := i
		for i < n && a[i] == a[j] {
			i++
		}
		arr = append(arr, pair{i - j, a[j]})
	}

	m := len(arr)
	pos := m
	sum := make([]int, m+1)

	next := make([]int, m)
	ans := []int{0, -1}

	for i := m - 1; i >= 0; i-- {
		sum[i] = sum[i+1] + arr[i].first
		if i+1 == m || arr[i+1].second > arr[i].second+1 {
			// 如果无法和后面的连续起来
			pos = i
		}
		next[i] = pos
		tmp := sum[i] - sum[pos+1]
		if tmp > ans[0] {
			ans[0] = tmp
			ans[1] = i
		}

		if arr[i].first == 1 {
			pos = i
		}
	}

	var first []int
	for i := 0; i < arr[ans[1]].first; i++ {
		first = append(first, arr[ans[1]].second)
	}

	var second []int
	for i := ans[1] + 1; i < next[ans[1]]; i++ {
		cur := arr[i]
		first = append(first, cur.second)
		for j := 0; j < cur.first-1; j++ {
			second = append(second, cur.second)
		}
	}
	// 中间的肯定至少2个
	if next[ans[1]] != ans[1] {
		for i := 0; i < arr[next[ans[1]]].first; i++ {
			first = append(first, arr[next[ans[1]]].second)
		}
	}

	reverse(second)

	first = append(first, second...)

	return first
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

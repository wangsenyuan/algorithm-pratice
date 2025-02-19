package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
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

func process(reader *bufio.Reader) (res []int) {
	n := readNum(reader)
	a, b := readTwoNums(reader)
	score := readNNums(reader, n)
	res = solve(a, b, score)
	return
}

func solve(a int, b int, score []int) []int {
	n := len(score)

	res := make([]int, n)

	if a == b {
		for i := 0; i < n; i++ {
			if i < a {
				res[i] = 1
			} else {
				res[i] = 2
			}
		}
		return res
	}
	type pair struct {
		first  int
		second int
	}

	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{score[i], i}
	}

	if a < b {
		slices.SortFunc(arr, func(x, y pair) int {
			if x.first != y.first {
				return y.first - x.first
			}
			// 相同的时候，优先使用前面的
			return x.second - y.second
		})
		for i := 0; i < a; i++ {
			res[arr[i].second] = 1
		}
		for i := a; i < n; i++ {
			res[arr[i].second] = 2
		}
	} else {
		// a > b
		slices.SortFunc(arr, func(x, y pair) int {
			if x.first != y.first {
				return y.first - x.first
			}
			return y.second - x.second
		})

		for i := 0; i < b; i++ {
			res[arr[i].second] = 2
		}
		for i := b; i < n; i++ {
			res[arr[i].second] = 1
		}
	}

	return res
}

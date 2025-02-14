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
	tc := readNum(reader)
	var buf bytes.Buffer
	for range tc {
		res := process(reader)
		for _, x := range res {
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

func process(reader *bufio.Reader) []int {
	n, m, q := readThreeNums(reader)
	a := readNNums(reader, q)
	return solve(n, m, a)
}

type pair struct {
	first  int
	second int
}

func solve(n int, m int, a []int) []int {
	// 可能的位置区间
	arr := []pair{{m - 1, m - 1}}

	play := func(i int) int {
		in := -1
		for j := 0; j < len(arr); j++ {
			if i < arr[j].first {
				arr[j].first--
			} else if arr[j].second < i {
				arr[j].second++
			} else {
				// i有可能就是joker的位置
				in = j
			}
		}
		if in >= 0 {
			arr = append(arr, pair{0, 0})
			arr = append(arr, pair{n - 1, n - 1})
			if arr[in].first == arr[in].second {
				copy(arr[in:], arr[in+1:])
				arr = arr[:len(arr)-1]
			}
		}

		slices.SortFunc(arr, func(x, y pair) int {
			return x.first - y.first
		})
		var sum int
		var k int
		for i := 0; i < len(arr); {
			l := arr[i].first
			r := arr[i].second
			for i < len(arr) && arr[i].first <= r+1 {
				// 有重叠的区域
				r = max(r, arr[i].second)
				i++
			}
			arr[k] = pair{l, r}
			k++
			sum += r - l + 1
		}
		arr = arr[:k]
		return sum
	}

	k := len(a)
	ans := make([]int, k)

	for i, x := range a {
		ans[i] = play(x - 1)
	}

	return ans
}

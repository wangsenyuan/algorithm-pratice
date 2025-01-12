package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)
	fmt.Println(res)
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

func solve(a []int) int {
	n := len(a)
	if n <= 2 {
		return 1
	}
	type pair struct {
		first  int
		second int
	}
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{a[i], i}
	}

	slices.SortFunc(arr, func(x, y pair) int {
		return x.first - y.first
	})
	// 如果删除0，删除1， 这两个需要特殊处理
	check := func(d int, deleted int) bool {
		prev := -1
		for i := range n {
			if i == deleted {
				continue
			}
			if prev >= 0 {
				if d != arr[i].first-arr[prev].first {
					return false
				}
			}
			prev = i
		}
		return true
	}

	if check(arr[2].first-arr[1].first, 0) {
		return arr[0].second + 1
	}
	if check(arr[2].first-arr[0].first, 1) {
		return arr[1].second + 1
	}
	d := arr[1].first - arr[0].first

	i := 1
	for i+1 < n && arr[i+1].first-arr[i].first == d {
		i++
	}

	if i == n-1 {
		// 这个和删除0是等价的，所以不会发生
		return arr[i].second + 1
	}
	if i == n-2 {
		return arr[n-1].second + 1
	}
	// i < n - 1
	j := i + 2
	if arr[j].first-arr[i].first != d {
		return -1
	}
	for j+1 < n && arr[j+1].first-arr[j].first == d {
		j++
	}
	if j == n-1 {
		return arr[i+1].second + 1
	}

	return -1
}

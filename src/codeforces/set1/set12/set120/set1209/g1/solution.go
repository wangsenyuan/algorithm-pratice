package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, _ := readTwoNums(reader)
	a := readNNums(reader, n)

	res := solve(a, nil)

	fmt.Println(res)
}

func solve(a []int, udpates [][]int) int {
	// len(updates) = 0
	first := make(map[int]int)
	last := make(map[int]int)

	for i, x := range a {
		if _, ok := first[x]; !ok {
			first[x] = i
		}
		last[x] = i
	}

	type pair struct {
		first  int
		second int
	}
	var arr []pair
	for x, i := range first {
		arr = append(arr, pair{i, last[x]})
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return a.first - b.first
	})

	var res int

	for i := 0; i < len(arr); {
		l := arr[i].first
		r := arr[i].second
		for i < len(arr) && arr[i].first <= r {
			r = max(r, arr[i].second)
			i++
		}
		freq := make(map[int]int)
		for j := l; j <= r; j++ {
			freq[a[j]]++
		}
		var v int
		for _, w := range freq {
			v = max(v, w)
		}
		res += (r - l + 1) - v
	}

	return res
}

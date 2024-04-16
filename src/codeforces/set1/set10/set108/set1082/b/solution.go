package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	readNum(reader)

	s := readString(reader)

	res := solve(s)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(s string) int {
	var arr []Pair
	n := len(s)
	var cnt int
	for i := 0; i <= n; i++ {
		if i == n || s[i] == 'S' {
			if cnt > 0 {
				arr = append(arr, Pair{i - cnt, i - 1})
			}
			cnt = 0
		} else {
			cnt++
		}
	}
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		// only one segments
		return arr[0].size()
	}
	if len(arr) == 2 {
		if arr[0].second+2 == arr[1].first {
			// can connect
			return arr[0].size() + arr[1].size()
		}
		return max(arr[0].size(), arr[1].size()) + 1
	}
	best := arr[0].size() + 1
	for i := 1; i < len(arr); i++ {
		best = max(best, arr[i].size()+1)
		if arr[i-1].second+2 == arr[i].first {
			// fill it
			best = max(best, arr[i-1].size()+arr[i].size()+1)
		}
	}

	return best
}

type Pair struct {
	first  int
	second int
}

func (p Pair) size() int {
	return p.second - p.first + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

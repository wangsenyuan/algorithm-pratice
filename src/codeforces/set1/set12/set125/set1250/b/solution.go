package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)

	a := readNNums(reader, n)

	res := solve(a, k)

	fmt.Println(res)
}

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

func solve(a []int, k int) int {
	freq := make([]int, k)
	for _, x := range a {
		freq[x-1]++
	}

	sort.Ints(freq)

	check := func(s int, r int) bool {
		if s < freq[k-1] {
			return false
		}
		var res int
		for l, r := 0, k-1; l <= r; r-- {
			res++
			if l < r && freq[l]+freq[r] <= s {
				l++
			}
		}
		return res <= r
	}

	get := func(r int) int {
		return sort.Search(2*freq[k-1], func(s int) bool {
			return check(s, r)
		})
	}

	best := math.MaxInt

	for i := max(1, (k+1)/2); i <= k; i++ {
		// 在这么多趟固定的情况下，如何找出最小的bus容量
		s := get(i)
		best = min(best, i*s)
	}

	return best
}

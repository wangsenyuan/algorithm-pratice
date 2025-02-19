package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func process(reader *bufio.Reader) int {
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	return solve(k, a, b)
}

type pair struct {
	first  int
	second int
}

func solve(k int, a []int, b []int) int {
	n := len(a)
	arr := make([]int, n*2)
	copy(arr, a)
	copy(arr[n:], b)

	sort.Ints(arr)
	var m int
	for i := 1; i <= len(arr); i++ {
		if i == len(arr) || arr[i] > arr[i-1] {
			arr[m] = arr[i-1]
			m++
		}
	}
	arr = arr[:m]
	bit := make(BIT, m+10)

	seq := make([]pair, n)
	for i := range n {
		j := sort.SearchInts(arr, a[i])
		seq[i] = pair{b[i], j}
		bit.update(j, 1)
	}

	slices.SortFunc(seq, func(x, y pair) int {
		return x.first - y.first
	})

	var ans int
	var i int
	for _, x := range arr {
		// cnt := bit.queryRange(0, cur.second)
		j := sort.SearchInts(arr, x)
		// arr[j] >= x
		cnt := bit.queryRange(0, j-1)
		if cnt <= k {
			ans = max(ans, x*(n-i))
		}
		for i < n && seq[i].first == x {
			bit.update(seq[i].second, -1)
			i++
		}
	}

	return ans
}

type BIT []int

func (bit BIT) update(p int, v int) {
	p++
	for p < len(bit) {
		bit[p] += v
		p += p & -p
	}
}

func (bit BIT) get(p int) int {
	p++
	var res int
	for p > 0 {
		res += bit[p]
		p -= p & -p
	}
	return res
}

func (bit BIT) queryRange(l int, r int) int {
	if l > r {
		return 0
	}
	res := bit.get(r)
	if l > 0 {
		res -= bit.get(l - 1)
	}
	return res
}

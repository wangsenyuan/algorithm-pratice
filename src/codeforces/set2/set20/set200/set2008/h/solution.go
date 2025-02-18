package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	q := make([]int, m)
	for i := range m {
		q[i] = readNum(reader)
	}
	return solve(a, q)
}

func solve(a []int, q []int) []int {
	n := len(a)
	freq := make([]int, n+1)
	for _, x := range a {
		freq[x]++
	}
	mem := make([]int, n+1)

	for i := 1; i <= n; i++ {
		freq[i] += freq[i-1]
		mem[i] = -1
	}

	check := func(x int, m int) bool {
		var cnt int
		for k := 0; k*x <= n && cnt*2 <= n; k++ {
			a := k * x
			b := min(k*x+m, n)
			cnt += freq[b]
			if a > 0 {
				cnt -= freq[a-1]
			}
		}

		return cnt*2 > n
	}

	get := func(x int) int {
		if mem[x] >= 0 {
			return mem[x]
		}

		l, r := 0, x-1
		for l < r {
			mid := (l + r) / 2
			if check(x, mid) {
				r = mid
			} else {
				l = mid + 1
			}
		}
		mem[x] = r
		return r
	}

	ans := make([]int, len(q))
	for i, x := range q {
		ans[i] = get(x)
	}

	return ans
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
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
	n, m, k := readThreeNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, m)
	return solve(a, b, k)
}

const inf = 1 << 60

func solve(a []int, b []int, k int) int {
	n := len(a)
	m := len(b)

	M := 1 << m
	vals := make([]int, M)
	for state := 0; state < M; state++ {
		v := 1<<30 - 1
		for i := 0; i < m; i++ {
			if (state>>i)&1 == 1 {
				v &= b[i]
			}
		}
		vals[state] = v
	}
	diffs := make([]int, m+1)
	var arr []int
	var sum int
	for _, num := range a {
		sum += num
		for i := 0; i <= m; i++ {
			diffs[i] = num
		}
		for state := 0; state < 1<<m; state++ {
			tmp := num & vals[state]
			c := bits.OnesCount(uint(state))
			diffs[c] = min(diffs[c], tmp)
		}
		for i := 1; i <= m; i++ {
			arr = append(arr, diffs[i-1]-diffs[i])
		}
	}

	sort.Ints(arr)

	for i := n*m - k; i < n*m; i++ {
		sum -= arr[i]
	}
	return sum
}

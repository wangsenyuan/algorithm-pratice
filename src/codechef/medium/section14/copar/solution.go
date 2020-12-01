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
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(n, A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

const MAX_A = 100010

var spf [MAX_A]int

func init() {
	set := make([]bool, MAX_A)
	for x := 2; x < MAX_A; x++ {
		if !set[x] {
			for y := x; y < MAX_A; y += x {
				if spf[y] == 0 {
					spf[y] = x
				}
				set[y] = true
			}
		}
	}
}

func solve(n int, A []int) int {
	left := make(map[int]int)
	right := make(map[int]int)

	update := func(mem map[int]int, k int, v int, fn func(int, int) int) {
		if w, ok := mem[k]; !ok {
			mem[k] = v
		} else {
			mem[k] = fn(w, v)
		}
	}

	for i := 0; i < n; i++ {
		cur := A[i]

		for spf[cur] != cur {
			x := spf[cur]
			update(left, x, i, min)
			update(right, x, i, max)
			cur /= x
		}
		update(left, cur, i, min)
		update(right, cur, i, max)
	}

	cnt := make([]int, n+1)

	for k, i := range left {
		j := right[k]
		cnt[i]++
		cnt[j]--
	}

	for i := 1; i <= n; i++ {
		cnt[i] += cnt[i-1]
	}
	for i := 1; i <= n; i++ {
		if cnt[i-1] == 0 {
			return i
		}
	}
	return -1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

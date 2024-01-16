package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(k, a)
		buf.WriteString(fmt.Sprintf("%d %d %d\n", res[0], res[1], res[2]))
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(k int, a []int) []int {
	// 对于j来说，就是要找和它尽量在高位相同的数
	// 排序后的前一个数？
	n := len(a)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		return a[id[i]] < a[id[j]]
	})

	res := []int{-1, -1, 0, -1}

	for i := 1; i < n; i++ {
		x := a[id[i-1]]
		y := a[id[i]]
		var tmp int
		for j := k - 1; j >= 0; j-- {
			if (x>>j)&1 != (y>>j)&1 {
				continue
			}
			if (x>>j)&1 == 0 {
				tmp ^= 1 << j
			}
		}
		if (x^tmp)&(y^tmp) > res[3] {
			res[0] = id[i-1] + 1
			res[1] = id[i] + 1
			res[2] = tmp
			res[3] = (x ^ tmp) & (y ^ tmp)
		}
	}
	return res[:3]
}

func solve1(k int, a []int) []int {
	// 是否能够在保证高位结果的情况下，让本位也是1

	best := []int{-1, -1, 0, 0}

	var dfs func(arr []int, d int, res int, xor int)

	dfs = func(arr []int, d int, res int, xor int) {
		if len(arr) <= 1 {
			return
		}
		if d < 0 {
			if xor >= best[3] {
				best[0] = arr[0] + 1
				best[1] = arr[1] + 1
				best[2] = res
				best[3] = xor
			}
			return
		}

		if best[3]>>(d+1) > xor>>(d+1) {
			return
		}

		var x, y []int
		for _, i := range arr {
			if (a[i]>>d)&1 == 0 {
				x = append(x, i)
			} else {
				y = append(y, i)
			}
		}
		if len(x) >= 2 {
			dfs(x, d-1, res|(1<<d), xor|(1<<d))
		}
		if len(y) >= 2 {
			dfs(y, d-1, res, xor|(1<<d))
		}
		dfs(arr, d-1, res, xor)
	}

	id := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		id[i] = i
	}

	dfs(id, k-1, 0, 0)

	return best[:3]
}

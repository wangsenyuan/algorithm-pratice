package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m, k := readThreeNums(reader)
	S := readNInt64s(reader, k)

	res := solve(n, m, S)

	var buf bytes.Buffer

	for _, ans := range res {
		for _, tmp := range ans {
			for i := 0; i < len(tmp); i++ {
				buf.WriteString(fmt.Sprintf("%d ", tmp[i]))
			}
			buf.WriteByte('\n')
		}
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

func solve(n int, m int, S []int64) [][][]int {
	// 从P开始， 选择K行，p * A[0] + (p + 1) * A[1] + .. (p + k - 1) * A[k-1] = sum
	// A[i] > 0 and A[i] <= m
	// 首先选取k
	// (p + k - 1) * m <= sum
	M := int64(m)
	N := int64(n)

	check := func(sum int64) [][]int {
		// p + k - 1
		// p + k - 1 <= n
		// 不需要这么复杂
		if sum <= N {
			//只需要选取一行， 选取那行就可以
			p := int(sum)
			k := 1
			return [][]int{
				{p, k},
				{1},
			}
		}
		p := 1
		k := n
		// sum > N
		// 所有行都放置1
		if sum < N*(N+1)/2 {
			// can put 1 at every row, p = 1, k = n
			k = search(n, func(i int) bool {
				return (2+int64(i)-1)*int64(i)/2 > sum
			})
			// k = 1 => 1 > sum => k > 1
			k--
		}

		L := make([]int, k)

		for i := 0; i < k; i++ {
			L[i]++
			sum -= int64(p + i)
		}

		// sum > N && sum < N * (N + 1) / 2
		// 无法给每行至少设置为1
		// 始终从p = 1, 开始, (p + p + k - 1) * k / 2 <= sum

		for i := k - 1; i >= 0; i-- {
			x := min(sum/int64(i+1), M-1)
			L[i] += int(x)
			sum -= x * int64(i+1)
		}

		return [][]int{
			{1, k},
			L,
		}
	}

	res := make([][][]int, len(S))

	for i, s := range S {
		res[i] = check(s)
	}

	return res
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(C []int) int {
	// | or &
	// 2 0 1 3 => 会变成2个数字
	// 固定一个边界i，一边取最小值，一边取最大值
	// 最小值部分用&， 最大值部分用｜
	// 那需要再确定一个边界j, or(i...j) and (i...j)
	// 从最高位考虑
	// or的部分，必须至少包含一个1，and的部分至少包含一个0
	// 然后找出所有的 10/01的节点
	// 但是这些节点然后怎么处理呢？
	// 高位假设按照 10分割了，那么当前位，需要找到那种01的节点
	n := len(C)
	N := 2 * n
	arr := make([]int, 2*N)

	for i := N; i < 2*N; i++ {
		arr[i] = C[i%n]
	}
	for i := N - 1; i > 0; i-- {
		arr[i] = arr[2*i] & arr[2*i+1]
	}

	const H = 31

	get := func(l int, r int) int {
		res := (1 << H) - 1
		l += N
		r += N
		for l < r {
			if l&1 == 1 {
				res = res & arr[l]
				l++
			}
			if r&1 == 1 {
				r--
				res = res & arr[r]
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	next := make([][]int, N)
	for i := N - 1; i >= 0; i-- {
		cur := C[i%n]
		next[i] = make([]int, H)
		for j := 0; j < H; j++ {
			if (cur>>j)&1 == 1 {
				next[i][j] = i
			} else if i+1 < N {
				next[i][j] = next[i+1][j]
			} else {
				next[i][j] = N
			}
		}
	}

	var res int

	for i := 0; i < n; i++ {
		m := i + n
		var or int
		ii := i
		for ii < m {
			or |= C[ii%n]
			and := get(ii+1, m)

			res = max(res, or-and)

			k := N
			for j := 0; j < H; j++ {
				if (or>>j)&1 == 0 {
					k = min(k, next[ii][j])
				}
			}
			ii = k
		}
	}

	return res
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

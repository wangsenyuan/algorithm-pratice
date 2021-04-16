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
		n, m, k := readThreeNums(reader)
		G := make([][]int, n)
		for i := 0; i < n; i++ {
			G[i] = readNNums(reader, m)
		}
		res := solve(G, k)
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

func solve(G [][]int, k int) int64 {
	n := len(G)
	m := len(G[0])
	S := make([][]int64, n)
	for i := 0; i < n; i++ {
		S[i] = make([]int64, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i > 0 {
				S[i][j] += S[i-1][j]
			}
			if j > 0 {
				S[i][j] += S[i][j-1]
			}
			if i > 0 && j > 0 {
				S[i][j] -= S[i-1][j-1]
			}
			S[i][j] += int64(G[i][j])
		}
	}

	K := int64(k)
	var res int64
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// find squars that starts at (i, j), n <= m
			l := search(min(n-i, m-j), func(l int) bool {
				tmp := S[i+l][j+l]
				if i > 0 {
					tmp -= S[i-1][j+l]
				}
				if j > 0 {
					tmp -= S[i+l][j-1]
				}
				if i > 0 && j > 0 {
					tmp += S[i-1][j-1]
				}
				return tmp >= K*int64(l+1)*int64(l+1)
			})
			//l++
			res += int64(min(n-i, m-j) - l)
		}
	}
	return res
}

func search(n int, fn func(int) bool) int {
	var left, right = 0, n
	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

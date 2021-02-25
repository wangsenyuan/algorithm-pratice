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
		grid := make([]string, n)
		for i := 0; i < n; i++ {
			grid[i], _ = reader.ReadString('\n')
			grid[i] = grid[i][:m]
		}
		res := solve(n, m, k, grid)
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

func solve(n, m, k int, grid []string) int {
	cnt := make([][]int, n)

	for i := 0; i < n; i++ {
		cnt[i] = make([]int, m)
		for j := 0; j < m; j++ {
			cnt[i][j] = int('1' - grid[i][j])
			if i > 0 {
				cnt[i][j] += cnt[i-1][j]
			}
			if j > 0 {
				cnt[i][j] += cnt[i][j-1]
			}
			if i > 0 && j > 0 {
				cnt[i][j] -= cnt[i-1][j-1]
			}
		}
	}

	tot := cnt[n-1][m-1]
	if tot == 0 || tot == n*m {
		return 0
	}

	check := func(w int) bool {
		// check whether a side w-length square got after (at most) k swaps
		for i := 0; i+w <= n; i++ {
			for j := 0; j+w <= m; j++ {
				a := cnt[i+w-1][j+w-1]
				if i > 0 {
					a -= cnt[i-1][j+w-1]
				}
				if j > 0 {
					a -= cnt[i+w-1][j-1]
				}
				if i > 0 && j > 0 {
					a += cnt[i-1][j-1]
				}
				b := w*w - a
				if b <= k && b <= tot-a {
					return true
				}
			}
		}
		return false
	}

	left, right := 1, min(n, m)+1
	for left < right {
		mid := (left + right) >> 1
		if !check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right - 1
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

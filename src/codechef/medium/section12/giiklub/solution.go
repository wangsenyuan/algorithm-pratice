package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		s, _ := reader.ReadBytes('\n')
		var n int
		pos := readInt(s, 0, &n)
		var X uint64
		readUint64(s, pos+1, &X)

		G := make([][]int64, n)
		for i := 0; i < n; i++ {
			G[i] = make([]int64, n)

			s, _ = reader.ReadBytes('\n')

			pos = 0
			var tmp uint64

			for j := 0; j < n; j++ {
				pos = readUint64(s, pos, &tmp) + 1
				G[i][j] = int64(tmp)
			}
		}

		fmt.Println(solve(n, int64(X), G))
	}
}

const MAX_N = 21

var C [MAX_N][MAX_N]int

func init() {
	C[0][0] = 1
	for i := 1; i < MAX_N; i++ {
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j] + C[i-1][j-1]
		}
	}
}

func solve(n int, X int64, G [][]int64) int64 {
	forward := make([][]int64, n)

	for i := 0; i < n; i++ {
		forward[i] = make([]int64, 0, C[n-1][i])
	}
	M := 1 << uint(n-1)

	for mask := M - 1; mask >= 0; mask-- {
		var x, y int
		var sum int64

		for i := 0; i < n-1; i++ {
			sum += G[x][y]
			if mask&(1<<uint(i)) > 0 {
				y++
			} else {
				x++
			}
		}
		forward[x] = append(forward[x], sum)
	}

	for i := 0; i < n; i++ {
		sort.Sort(Int64Array(forward[i]))
	}

	var res int64

	for mask := M - 1; mask >= 0; mask-- {
		x, y := n-1, n-1
		var sum int64
		for i := 0; i < n-1; i++ {
			sum += G[x][y]
			if mask&(1<<uint(i)) > 0 {
				y--
			} else {
				x--
			}
		}

		sum += G[x][y]

		if sum <= X {
			res += countSmaller(forward[x], X-sum)
		}
	}
	return res
}

func countSmaller(arr []int64, X int64) int64 {
	var left int
	var right = len(arr)

	for left < right {
		mid := (left + right) / 2
		if arr[mid] > X {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return int64(right)
}

type Int64Array []int64

func (arr Int64Array) Len() int {
	return len(arr)
}

func (arr Int64Array) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr Int64Array) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

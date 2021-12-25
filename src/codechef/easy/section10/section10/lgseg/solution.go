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
		n, k, s := readThreeNums(reader)
		A := readNNums(reader, n)
		res := solve(n, k, s, A)
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

const H = 20

func solve(n int, k int, s int, A []int) int {
	sum := make([]int64, n)
	for i := 0; i < n; i++ {
		sum[i] = int64(A[i])
		if i > 0 {
			sum[i] += sum[i-1]
		}
	}

	next := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		next[i] = make([]int, H)
	}

	for j := 0; j < H; j++ {
		next[n][j] = n
	}

	S := int64(s)

	for i := n - 1; i >= 0; i-- {
		var prev int64
		if i > 0 {
			prev = sum[i-1]
		}
		j := sort.Search(n, func(j int) bool {
			return sum[j]-prev > S
		})
		next[i][0] = j

		for h := 1; h < H; h++ {
			next[i][h] = next[next[i][h-1]][h-1]
		}
	}

	var best int

	for i := 0; i < n; i++ {
		tmp := k
		u := i
		for j := H - 1; j >= 0; j-- {
			if tmp >= (1 << uint(j)) {
				u = next[u][j]
				tmp -= 1 << uint(j)
			}
		}
		best = max(best, u-i)
	}

	return best
}


func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

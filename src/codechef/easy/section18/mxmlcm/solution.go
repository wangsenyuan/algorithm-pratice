package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		fmt.Println(solve(n, m, A))
	}
}

const MAX_N = 10000

var PF []int

func init() {
	PF = make([]int, MAX_N+1)

	for i := 2; i <= MAX_N; i++ {
		if PF[i] == 0 {
			for j := i; j <= MAX_N; j += i {
				if PF[j] == 0 {
					PF[j] = i
				}
			}
		}
	}
}

func solve(n, m int, A []int) int {
	F := make([]int, MAX_N+1)

	for i := 0; i < n; i++ {
		x := A[i]
		for x > 1 {
			y := PF[x]
			var cnt int
			for x%y == 0 {
				cnt++
				x /= y
			}
			F[y] = max(F[y], cnt)
		}
	}

	var best = 1
	var res = 1

	for i := 2; i <= m; i++ {
		gain := 1
		x := i
		for x > 1 {
			y := PF[x]
			var cnt int
			for x%y == 0 {
				cnt++
				x /= y
			}
			cnt -= F[y]
			for cnt > 0 {
				gain *= y
				cnt--
			}
		}
		if gain > best {
			best = gain
			res = i
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

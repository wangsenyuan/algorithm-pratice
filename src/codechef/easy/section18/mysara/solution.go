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
		n := readNum(reader)
		B := readNNums(reader, n)
		fmt.Println(solve(n, B))
	}
}

const MOD = 1000000007

func solve(n int, B []int) int {
	set := make([]bool, 31)
	var cnt int
	for i := 0; i < 31; i++ {
		if B[0]&(1<<uint(i)) > 0 {
			set[i] = true
			cnt++
		}
	}
	X := B[0]

	var res int64 = 1

	for i := 1; i < n; i++ {
		x := B[i]

		if x&X != X {
			return 0
		}

		X |= x

		tmp := 1 << uint(cnt)

		res *= int64(tmp)
		res %= MOD

		for j := 0; j < 31; j++ {
			if x&(1<<uint(j)) > 0 {
				if set[j] {
					continue
				}
				set[j] = true
				cnt++
			}
		}
	}

	return int(res)
}

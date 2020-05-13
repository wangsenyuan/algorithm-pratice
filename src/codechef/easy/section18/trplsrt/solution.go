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
		n, k := readTwoNums(reader)
		A := readNNums(reader, n)
		ok, res := solve(n, k, A)
		if !ok {
			fmt.Println("-1")
			continue
		}
		fmt.Println(len(res))
		for _, arr := range res {
			fmt.Printf("%d %d %d\n", arr[0], arr[1], arr[2])
		}
	}
}

func solve(N, K int, A []int) (bool, [][]int) {
	B := make([]int, N+1)
	copy(B[1:], A)

	cycles := make([][]Pair, 0, N)
	oddCycles := make([][]Pair, 0, N)

	tmp := make([]Pair, 0, N)

	var parity int

	for i := 1; i <= N; i++ {
		for i != B[i] {
			tmp = append(tmp, Pair{i, B[i]})
			parity++
			B[i], B[B[i]] = B[B[i]], B[i]
		}
		if len(tmp) == 0 {
			continue
		}
		arr := make([]Pair, len(tmp))
		copy(arr, tmp)
		tmp = tmp[:0]
		if len(tmp)%2 == 1 {
			oddCycles = append(oddCycles, arr)
		} else {
			cycles = append(cycles, arr)
		}
	}

	if parity%2 == 1 {
		return false, nil
	}
	// put odd cycles at end of cycles
	for _, arr := range oddCycles {
		cycles = append(cycles, arr)
	}

	swaps := make([]Pair, 0, parity)

	for _, arr := range cycles {
		for _, sp := range arr {
			if sp.first > sp.second {
				swaps = append(swaps, Pair{sp.second, sp.first})
			} else {
				swaps = append(swaps, sp)
			}
		}
	}
	res := make([][]int, 0, len(swaps))
	for i := 0; i < len(swaps); i += 2 {
		f, g := swaps[i], swaps[i+1]

		i1, j1 := f.first, f.second
		i2, j2 := g.first, g.second

		if i1 == i2 {
			res = append(res, []int{i1, j1, j2})
			continue
		}
		if j1 == j2 {
			res = append(res, []int{j1, i1, i2})
			continue
		}

		if i1 == j2 {
			res = append(res, []int{i1, j1, i2})
			continue
		}

		if i2 == j1 {
			res = append(res, []int{i2, i1, j2})
			continue
		}
		res = append(res, []int{i1, j1, j2})
		res = append(res, []int{i2, j2, i1})
	}

	return true, res
}

type Pair struct {
	first, second int
}

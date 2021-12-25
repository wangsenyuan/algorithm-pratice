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
		A := make([]uint64, n)
		s, _ := reader.ReadBytes('\n')
		var pos int
		for i := 0; i < n; i++ {
			pos = readUint64(s, pos, &A[i]) + 1
		}
		res := solve(n, A)
		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func solve(n int, A []uint64) bool {
	if n == 1 {
		return true
	}

	all := make(map[uint64]bool)

	mem := make(map[uint64]bool)

	mem[A[0]] = true
	all[A[0]] = true

	for i := 1; i < n; i++ {
		if mem[A[i]] {
			return false
		}
		newMem := make(map[uint64]bool)
		for x := range mem {
			y := x | A[i]
			if all[y] || newMem[y] {
				return false
			}
			newMem[y] = true
		}

		if newMem[A[i]] {
			return false
		}

		newMem[A[i]] = true

		for x := range newMem {
			all[x] = true
		}

		mem = newMem
	}

	return true
}

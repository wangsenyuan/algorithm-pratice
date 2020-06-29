package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, X := readTwoNums(reader)
		A := readNNums(reader, n)
		fmt.Println(solve(n, A, X))
	}
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

func solve(N int, A []int, X int) int64 {
	var res int64

	freq := make([]int, X+1)
	sum := make([]int, N)

	process := func(sz int) {
		if sz > N {
			return
		}
		var i int
		var tot int64
		for i < sz-1 {
			tot += int64(A[i])
			i++
		}

		S := X / sz

		var p int

		for i < N {
			tot += int64(A[i])
			if tot <= int64(S) {
				sum[p] = int(tot)
				p++
				freq[int(tot)]++
			}
			tot -= int64(A[i-sz+1])
			i++
		}
		for i := 0; i < p; i++ {
			res += int64(freq[S-sum[i]])
		}

		for i := 0; i < p; i++ {
			freq[sum[i]]--
		}
	}

	for x := 1; x*x <= X; x++ {
		if X%x != 0 {
			continue
		}
		process(x)
		y := X / x
		if y != x {
			process(y)
		}
	}

	return res
}

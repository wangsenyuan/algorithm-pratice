package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	D := readNNums(reader, n)
	ask := func(t, u int) []int {
		fmt.Printf("%d %d\n", t, u)
		n := readNum(reader)
		return readNNums(reader, n)
	}

	solve(n, k, D, ask)

	fmt.Print("3\n")
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

func solve(n, k int, D []int, ask func(t int, u int) []int) {
	var ii int

	for i := 1; i < n; i++ {
		if D[i] > D[ii] {
			ii = i
		}
	}

	mark := make([]bool, n)
	if D[ii] >= k {
		mark[ii] = true
		res := ask(2, ii+1)
		for _, i := range res {
			mark[i-1] = true
			tmp := ask(1, i)
			for _, j := range tmp {
				mark[j-1] = true
			}
		}
	}

	for i := 0; i < n; i++ {
		if mark[i] {
			continue
		}
		mark[i] = true
		res := ask(1, i+1)
		for _, j := range res {
			mark[j-1] = true
		}
	}
}

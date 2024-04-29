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

	n := readNum(reader)

	A := readNNums(reader, n)
	B := readNNums(reader, n)

	fmt.Println(solve(n, A, B))
}

func solve(n int, A []int, B []int) int {
	pos := make([]int, n+1)
	for i, num := range B {
		pos[num] = i
	}
	cnt := make([]int, n+1)
	for i := 0; i < n; i++ {
		j := pos[A[i]]
		cnt[(i-j+n)%n]++
	}

	res := cnt[0]
	for i := 1; i <= n; i++ {
		res = max(res, cnt[i])
	}
	return res
}

func solve1(n int, A []int, B []int) int {
	pos := make([]int, n+1)

	for i := 0; i < n; i++ {
		pos[A[i]] = i
	}
	left := make([]int, n)
	right := make([]int, n)

	for i := 0; i < n; i++ {
		j := pos[B[i]]
		if i == j {
			left[0]++
			continue
		}
		if j < i {
			left[i-j]++
			right[j-i+n]++
		} else {
			left[i-j+n]++
			right[j-i]++
		}
	}

	res := left[0]

	for k := 1; k < n; k++ {
		if left[k] > res {
			res = left[k]
		}
		if right[k] > res {
			res = right[k]
		}
	}
	return res
}

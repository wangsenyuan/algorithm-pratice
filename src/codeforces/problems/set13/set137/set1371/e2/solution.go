package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, p := readTwoNums(reader)
	A := readNNums(reader, n)
	res := solve(n, p, A)
	fmt.Println(len(res))
	for i := 0; i < len(res); i++ {
		fmt.Printf("%d ", res[i])
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

const MAX_N = 262144

var BK [MAX_N]int

func solve(n int, p int, A []int) []int {
	var bas = A[0]

	for i := 1; i < n; i++ {
		bas = max(bas, A[i])
	}

	for i := 0; i < MAX_N; i++ {
		BK[i] = 0
	}

	for i := 0; i < n; i++ {
		BK[max(0, A[i]-bas+n)]++
	}

	for i := 1; i < MAX_N; i++ {
		BK[i] += BK[i-1]
	}

	var st = 1
	for i := 1; i <= n; i++ {
		for BK[st+i-1] < i {
			st++
		}
	}

	st += bas - n

	fi := n

	for i := 1; i <= n; i++ {
		for fi > 0 {
			if BK[fi+(i-1)]-(i-1) < p {
				break
			}
			fi--
		}
	}

	fi += bas - n

	if fi < st {
		return nil
	}

	res := make([]int, 0, fi-st+1)

	for i := st; i <= fi; i++ {
		res = append(res, i)
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

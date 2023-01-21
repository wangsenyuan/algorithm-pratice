package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func solve(n int, p int, A []int) []int {
	sort.Ints(A)
	X := A[n-1]

	// x < X - n + 1 f(x) == 0
	// x >= X, f(x) = n!, and p < n
	var res []int
	var jj int
	for x := X - n + 1; x < X; x++ {
		for jj < n && A[jj] < x {
			jj++
		}
		j := jj
		var ok = true
		v := j
		for i := x; i <= x+n-1; i++ {
			k := j
			for k < n && A[k] == i {
				k++
			}
			v += k - j
			if v%p == 0 {
				ok = false
				break
			}
			v--
			j = k
		}
		if ok {
			res = append(res, x)
		}
	}
	return res
}

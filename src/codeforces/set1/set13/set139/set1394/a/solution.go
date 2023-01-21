package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, d, m := readThreeNums(reader)
	A := readNNums(reader, n)
	fmt.Println(solve(n, d, m, A))
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

const N_INF = -(int64(1) << 60)

func solve(n int, d int, m int, A []int) int64 {
	X := make([]int, 0, n)
	Y := make([]int, 0, n)

	for i := 0; i < n; i++ {
		if A[i] > m {
			X = append(X, A[i])
		} else {
			Y = append(Y, A[i])
		}
	}

	sort.Ints(X)
	sort.Ints(Y)

	reverse(X)
	reverse(Y)

	XX := sum(X)
	YY := sum(Y)

	if len(X) == 0 {
		return YY[n]
	}

	l := len(YY)
	for j := l; j <= n; j++ {
		YY = append(YY, YY[l-1])
	}

	k := len(X)
	var res int64
	N := int64(n)
	for i := (k + d) / (d + 1); i <= k; i++ {
		if int64(i-1)*int64(d+1)+1 <= N {
			j := n - (i-1)*(d+1) - 1
			tmp := XX[i] + YY[j]
			if tmp > res {
				res = tmp
			}
		}
	}
	return res
}

func sum(arr []int) []int64 {
	n := len(arr)
	res := make([]int64, n+1)
	for i := 0; i < n; i++ {
		res[i+1] = res[i] + int64(arr[i])
	}
	return res
}

type Pair struct {
	first  int
	second int64
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

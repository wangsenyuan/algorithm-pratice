package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		cnt, arr := solve(n)
		buf.WriteString(fmt.Sprintf("%d\n", cnt))
		for i := 0; i < len(arr); i++ {
			buf.WriteString(fmt.Sprintf("%d ", arr[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

const N = 18

const MOD = 1000000007

var P []int64

var PW []int64

func init() {
	T := 1 << uint(N)
	// abount 1000000
	P = make([]int64, T)
	P[0] = 1
	for i := 1; i < T; i++ {
		P[i] = int64(i) * P[i-1] % MOD
	}

	PW = make([]int64, N)
	PW[0] = 1
	for i := 1; i < N; i++ {
		PW[i] = 2 * PW[i-1] % MOD
	}
}

func solve(n int) (int, []int) {
	if n == 1 {
		return 2, []int{0, 1}
	}

	var res int64 = 4 * P[PW[n-1]] % MOD

	arr := make([]int, 1<<uint(n))

	dir := 1

	X := 1 << uint(n-1)
	var j int
	for i := 0; i < X*2; i += 2 {
		if dir == 1 {
			arr[i] = j
			arr[i+1] = j ^ X
		} else {
			arr[i] = j ^ X
			arr[i+1] = j
		}
		j++
		dir *= -1
	}

	return int(res), arr
}

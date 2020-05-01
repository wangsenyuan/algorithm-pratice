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
		N, K := readTwoNums(reader)

		A := readNNums(reader, N)

		fmt.Println(solve(N, K, A))
	}
}

func solve(N int, K int, A []int) int {
	// double with 2, to avoid divide 2 get non-integer points
	// K *= 2

	h := N >> 1
	// var R int

	// ps := make([]Pair, h)
	cnt := make([]int, 2*K+1)

	for i := 0; i < h; i++ {
		cnt[A[i]+A[N-1-i]]++
	}

	pref := make([]int, 2*K+2)

	for i := 0; i < h; i++ {
		a, b := A[i], A[N-1-i]
		if a > b {
			a, b = b, a
		}
		// a < b
		pref[a+1]++
		pref[b+K+1]--
	}

	for i := 1; i < len(pref); i++ {
		pref[i] += pref[i-1]
	}

	best := N

	for sum := 2; sum <= 2*K; sum++ {
		temp := pref[sum] - cnt[sum] + (h-pref[sum])*2
		if temp < best {
			best = temp
		}
	}

	return best
}

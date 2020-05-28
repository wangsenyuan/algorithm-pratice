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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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
		H := readNum(reader)
		N, M := readTwoNums(reader)
		line := readNNums(reader, 3)
		A, B, I := line[0], line[1], line[2]
		D := readNNums(reader, M)
		res := solve(H, N, M, A, B, I, D)
		var r int
		if res&1 == 1 {
			r = 5
		}
		fmt.Printf("%d.%d\n", res/2, r)
	}
}

func solve(H, N, M, A, B, IND int, D []int) int64 {
	// result is always W * H / 2
	// W = X[N-1]
	pos := make([]int, M)

	S := make([]int64, M)

	S[IND] = 0

	var i int = 1

	var X int64
	for pos[IND] == 0 && i < N {
		pos[IND] = i
		i++
		X += int64(D[IND])
		IND = int((int64(A)*int64(IND) + int64(B)) % int64(M))
		if pos[IND] == 0 {
			S[IND] = X
		}
	}
	if i < N {
		j, r := (N-pos[IND])/(i-pos[IND]), (N-pos[IND])%(i-pos[IND])

		prev := S[IND]

		W := (X-prev)*int64(j) + prev

		for r > 0 {
			W += int64(D[IND])
			IND = int((int64(A)*int64(IND) + int64(B)) % int64(M))
			r--
		}
		X = W
	}

	return int64(H) * X
}

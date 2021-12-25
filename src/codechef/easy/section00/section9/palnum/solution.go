package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		N, M, K := readThreeNums(reader)
		S, _ := reader.ReadString('\n')
		ops := make([][]int, M)
		for i := 0; i < M; i++ {
			ops[i] = readNNums(reader, 3)
		}
		ok, res := solve(N, M, K, S, ops)

		if !ok {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(res)
			buf.WriteByte('\n')
		}
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

func solve(N int, M int, k int, S string, ops [][]int) (bool, string) {
	K := int64(k)
	cost := make([][]int64, 10)

	for i := 0; i < 10; i++ {
		cost[i] = make([]int64, 10)
		for j := 0; j < 10; j++ {
			cost[i][j] = K + 1
		}
		cost[i][i] = 0
	}

	for i := 0; i < M; i++ {
		op := ops[i]
		x, y, w := op[0], op[1], op[2]
		cost[x][y] = int64(w)
	}

	for j := 0; j < 10; j++ {
		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				cost[x][y] = min(cost[x][y], cost[x][j]+cost[j][y])
			}
		}
	}

	half := N / 2
	X := make([]int64, half+(N&1))
	P := make([]int, half+(N&1))
	var took int64
	for i := 0; i < half && took <= K; i++ {
		x := int(S[i] - '0')
		y := int(S[N-1-i] - '0')
		// let's find the min cost to make it palindrome first
		pick := -1
		score := K + 1
		for j := 0; j < 10; j++ {
			if cost[x][j]+cost[y][j] < score {
				pick = j
				score = cost[x][j] + cost[y][j]
			}
		}
		if pick < 0 {
			return false, ""
		}
		X[i] = score
		P[i] = pick
		took += score
	}
	if N&1 == 1 {
		P[half] = int(S[half] - '0')
	}
	if took > K {
		return false, ""
	}
	// took <= K and it is palindrome now
	for i := 0; i < half; i++ {
		x := int(S[i] - '0')
		y := int(S[N-1-i] - '0')
		// we want to pick 9, 8, 7...
		tmp := took - X[i]
		for j := 9; j > P[i]; j-- {
			if tmp+cost[x][j]+cost[y][j] <= K {
				// good to replace P[i] by j
				P[i] = j
				took = tmp + cost[x][j] + cost[y][j]
				break
			}
		}
	}
	if N&1 == 1 {
		x := int(S[half] - '0')
		for j := 9; j > x; j-- {
			if took+cost[x][j] <= K {
				P[half] = j
				break
			}
		}
	}

	buf := make([]byte, N)
	for i := 0; i < half; i++ {
		buf[i] = byte('0' + P[i])
		buf[N-1-i] = byte('0' + P[i])
	}
	if N&1 == 1 {
		buf[half] = byte('0' + P[half])
	}

	return true, string(buf)
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

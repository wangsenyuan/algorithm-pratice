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
		S, _ := reader.ReadString('\n')
		T, _ := reader.ReadString('\n')
		res := solve(S[:len(S)-1], T[:len(T)-1])

		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
}

const N = 100100

const H = 520

var jump [N][H]int

var W [H]string

var SS [N + H + 100]byte

var F [N]int

var ans [N]int

func solve(S, T string) []int {
	var wn int
	for i, j := 0, 0; i <= len(S); i++ {
		if i == len(S) || S[i] == '*' {
			if j > 0 {
				wn++
				W[wn] = string(SS[:j])
				j = 0
			}
		} else {
			SS[j] = S[i]
			j++
		}
	}

	for i := 1; i <= wn; i++ {
		SS[0] = '#'
		copy(SS[1:], W[i])
		SS[1+len(W[i])] = '#'
		copy(SS[2+len(W[i]):], T)
		ssn := 2 + len(W[i]) + len(T)
		F[1] = 0
		for j := 2; j < ssn; j++ {
			k := F[j-1]
			for k > 0 && SS[j] != SS[k+1] {
				k = F[k]
			}
			F[j] = k
			if SS[j] == SS[k+1] {
				F[j]++
			}
		}
		jump[i][len(T)+10] = len(T) + 10
		jump[i][len(T)] = len(T) + 10
		for j := 0; j < len(W[i]); j++ {
			jump[i][len(T)-j] = len(T) + 10
		}
		curPos := len(T) - len(W[i])
		for j := ssn - 1; j >= len(W[i])+2 && curPos >= 0; j, curPos = j-1, curPos-1 {
			if F[j] != len(W[i]) {
				jump[i][curPos] = jump[i][curPos+1]
			} else {
				jump[i][curPos] = curPos + len(W[i])
			}
		}
	}

	for i := 0; i < len(T); i++ {
		ps := i

		for j := 1; j <= wn; j++ {
			ps = jump[j][ps]
		}

		if ps > len(T) {
			ans[i] = -1
		} else {
			ans[i] = ps
		}
		if wn == 0 {
			ans[i] = i + 1
		}
	}

	return ans[:len(T)]
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

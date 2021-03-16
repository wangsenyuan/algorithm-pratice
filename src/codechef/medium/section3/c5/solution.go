package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	A := readNNums(reader, n)
	m := readNum(reader)
	B := readNNums(reader, m)
	res := solve(A, B)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	if len(res) > 0 {
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
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

func solve(A []int, B []int) []int {
	L := make([]int, len(B))

	P := make([][]int, len(A))
	PL := make([][]int, len(A))

	for i := 0; i < len(A); i++ {
		P[i] = make([]int, len(B))
		PL[i] = make([]int, len(B))
	}

	var ii, jj int
	var ans int

	for i := 0; i < len(A); i++ {
		var cur int
		maxJ := -1

		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				if cur+1 > L[j] {
					L[j] = cur + 1
					P[i][j] = maxJ
					PL[i][j] = L[j]
				}
				if L[j] > ans {
					ans = L[j]
					ii = i
					jj = j
				}
			} else {
				if A[i] > B[j] && cur < L[j] {
					cur = L[j]
					maxJ = j
				}
			}
		}
	}

	var print func(row int, col int, ans int)
	res := make([]int, 0, ans)

	print = func(row int, col int, ans int) {
		if col < 0 || ans == 0 {
			return
		}
		if A[row] == B[col] && PL[row][col] == ans {
			print(row-1, P[row][col], ans-1)
			res = append(res, A[row])
			return
		}
		print(row-1, col, ans)
	}

	if ans == 0 {
		return nil
	}

	print(ii-1, P[ii][jj], ans-1)

	res = append(res, A[ii])

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

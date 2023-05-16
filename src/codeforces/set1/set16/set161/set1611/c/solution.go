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
		//readString(reader)
		n := readNum(reader)
		A := readNNums(reader, n)

		res := solve(A)

		if len(res) == 0 {
			buf.WriteString("-1\n")
			continue
		}

		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(A []int) []int {
	n := len(A)
	if n == 1 {
		return A
	}
	// P[l] < P[r]  then P[l] 加到A的左边
	// P[l] > P[r]  then P[r] 加到A的右边
	//   对于第二种情况，假设 P[l] 直到越到一个更大的 P[r'], 这时 P[l]加到了左边
	// n要么在第一位 A[0] = n, 或者 A[n-1] = n
	if A[0] != n && A[n-1] != n {
		return nil
	}
	// 如果 A[n-1] = n
	// n-1在遇到n之前，是不会被放入A的，假设pos[n-1] = i, 且 pos[n] = j, 如果 i < j
	// 那么在j后面的都会放在n-1的后面，且顺序翻转
	// i...j在左边，且顺序翻转
	// 也就是在A中，在n-1前面的，顺序要翻过来，就是它们在P中的顺序，且占连续的区间
	//   n-1后面的部分，也翻过来即可
	P := make([]int, n)
	if A[0] == n {
		i := 1
		for i < n && A[i] != n-1 {
			i++
		}
		copy(P, A[1:i+1])
		reverse(P[0:i])
		P[i] = n
		if i+1 < n {
			copy(P[i+1:], A[i+1:])
			reverse(P[i+1:])
		}
	} else {
		i := 0
		for i < n && A[i] != n-1 {
			i++
		}
		copy(P, A[0:i+1])
		reverse(P[0 : i+1])
		P[i+1] = n
		if i+2 < n {
			copy(P[i+2:], A[i+1:])
			reverse(P[i+2:])
		}
	}

	return P
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

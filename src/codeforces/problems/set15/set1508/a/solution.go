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
		n := readNum(reader)
		A, _ := reader.ReadString('\n')
		B, _ := reader.ReadString('\n')
		C, _ := reader.ReadString('\n')
		res := solve(n, A, B, C)
		buf.WriteString(res)
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

func solve(n int, A, B, C string) string {
	v0 := make([]string, 0, 2)
	v1 := make([]string, 0, 2)
	if label(n, A) {
		v0 = append(v0, A)
	} else {
		v1 = append(v1, A)
	}
	if label(n, B) {
		v0 = append(v0, B)
	} else {
		v1 = append(v1, B)
	}
	if label(n, C) {
		v0 = append(v0, C)
	} else {
		v1 = append(v1, C)
	}

	if len(v0) > 1 {
		return mix(n, v0[0], v0[1], '0')
	}
	return mix(n, v1[0], v1[1], '1')
}

func mix(n int, a, b string, c byte) string {
	v1 := make([]int, 0, n)
	v2 := make([]int, 0, n)
	v1 = append(v1, -1)
	v2 = append(v2, -1)

	for i := 0; i < 2*n; i++ {
		if a[i] == c {
			v1 = append(v1, i)
		}
		if b[i] == c {
			v2 = append(v2, i)
		}
	}

	var res bytes.Buffer

	for i := 0; i < n; i++ {
		for j := v1[i] + 1; j < v1[i+1]; j++ {
			res.WriteByte(a[j])
		}
		for j := v2[i] + 1; j < v2[i+1]; j++ {
			res.WriteByte(b[j])
		}
		res.WriteByte(c)
	}

	for j := v1[n] + 1; j < 2*n; j++ {
		res.WriteByte(a[j])
	}
	for j := v2[n] + 1; j < 2*n; j++ {
		res.WriteByte(b[j])
	}

	return res.String()
}

func label(n int, s string) bool {
	var cnt0 int
	for i := 0; i < 2*n; i++ {
		if s[i] == '0' {
			cnt0++
		}
	}
	return cnt0 >= n
}

func solve1(n int, A, B, C string) string {
	var p1, p2, p3 int
	var buf bytes.Buffer

	N := 2 * n
	for p1 < N && p2 < N && p3 < N {
		if A[p1] == B[p2] {
			buf.WriteByte(A[p1])
			p1++
			p2++
		} else if A[p1] == C[p3] {
			buf.WriteByte(A[p1])
			p1++
			p3++
		} else {
			// B[p2] == C[p3]
			buf.WriteByte(B[p2])
			p2++
			p3++
		}
	}
	l := buf.Len() - n
	if p1 == N {
		if p2 >= l {
			buf.WriteString(B[p2:N])
		} else {
			buf.WriteString(C[p3:N])
		}
	} else if p2 == N {
		if p1 >= l {
			buf.WriteString(A[p1:N])
		} else {
			buf.WriteString(C[p3:N])
		}
	} else {
		// p3 == N
		if p1 >= l {
			buf.WriteString(A[p1:N])
		} else {
			buf.WriteString(B[p2:N])
		}
	}

	return buf.String()
}

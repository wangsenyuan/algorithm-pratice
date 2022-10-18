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
		A := readNInt64s(reader, n)
		ok, res := solve(A)
		if !ok {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", len(res)))
			for _, op := range res {
				buf.WriteString(fmt.Sprintf("%d %d\n", op[0], op[1]))
			}
		}
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(A []int64) (bool, [][]int) {
	// if A is already good, then 0
	// A[0....n-1] and A[1...n], must has gcd > 1
	// gcd(A[0], A[n-1]) = 1

	n := len(A)

	pref := make([]int64, n)
	pref[0] = A[0]
	for i := 1; i < n; i++ {
		pref[i] = gcd(pref[i-1], A[i])
	}
	suf := make([]int64, n)
	suf[n-1] = A[n-1]
	for i := n - 2; i >= 0; i-- {
		suf[i] = gcd(suf[i+1], A[i])
	}
	ans := 2
	var found int
	var first int
	if suf[1] != 1 {
		ans--
		found++
		first++
	}
	if pref[n-2] != 1 {
		ans--
		found++
	}

	var op []int

	for i := 1; found < 2 && i+1 < n; i++ {
		if gcd(pref[i-1], suf[i+1]) > 1 {
			op = append(op, i+1)
			found++
		}
	}
	if found < 2 {
		return false, nil
	}

	var res [][]int

	if ans == 1 {
		if first > 0 {
			res = append(res, []int{op[0], n})
		} else {
			res = append(res, []int{1, op[0]})
		}
	} else if ans == 2 {
		res = append(res, []int{1, op[0]})
		res = append(res, []int{op[1], n})
	}

	return true, res
}

func gcd_of_array(arr []int64) int64 {
	var res int64

	for _, num := range arr {
		res = gcd(res, num)
	}
	return res
}

func reverse(arr []int64) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

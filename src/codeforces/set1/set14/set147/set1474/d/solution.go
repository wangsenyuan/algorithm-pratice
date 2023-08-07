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
		A := readNNums(reader, n)
		res := solve(A)

		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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

func solve(A []int) bool {
	n := len(A)
	pref := make([]int, n+2)
	suf := make([]int, n+2)
	for i := 0; i < n+2; i++ {
		pref[i] = -1
		suf[i] = -1
	}
	pref[0] = 0

	for i := 1; i <= n; i++ {
		a := pref[i-1]
		cur := A[i-1] - a
		if cur < 0 {
			break
		}
		pref[i] = cur
	}

	if pref[n] == 0 {
		return true
	}

	suf[n+1] = 0

	for i := n; i >= 1; i-- {
		a := suf[i+1]
		cur := A[i-1] - a
		if cur < 0 {
			break
		}
		suf[i] = cur
	}

	if suf[1] == 0 {
		return true
	}

	check := func(a, b, c, d int) bool {
		if a < 0 || d < 0 {
			return false
		}
		// swap
		b, c = c, b
		if b < a || c < d {
			return false
		}
		return b-a == c-d
	}

	for i := 1; i < n; i++ {
		if check(pref[i-1], A[i-1], A[i], suf[i+2]) {
			return true
		}
	}

	return false
}

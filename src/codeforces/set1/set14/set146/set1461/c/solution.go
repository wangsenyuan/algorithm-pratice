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
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		exp := make([]string, m)
		for i := 0; i < m; i++ {
			exp[i] = readString(reader)
		}
		res := solve(A, exp)
		buf.WriteString(fmt.Sprintf("%.8f\n", res))
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

func solve(A []int, exp []string) float64 {
	n := len(A)

	// 如果 (r, p)
	// dp[i] 表示到i时，前缀不sorted的概率

	R := make([]int, len(exp))
	P := make([]float64, len(exp))

	for i := 0; i < len(exp); i++ {
		R[i], P[i] = format(exp[i])
	}

	correct := n - 1
	for correct >= 0 && A[correct] == correct+1 {
		correct--
	}

	ans := 1.0

	if correct < 0 {
		ans = 0.0
	}

	for i := 0; i < len(R); i++ {
		r := R[i] - 1
		if r >= correct {
			ans = ans * (1 - P[i])
		}
	}

	return 1.0 - ans
}

func format(s string) (int, float64) {
	buf := []byte(s)
	var i int
	var p float64
	pos := readInt(buf, 0, &i)
	readFloat64(buf, pos+1, &p)

	return i, p
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

func readFloat64(bs []byte, from int, val *float64) int {
	i := from
	sign := 1
	if bs[i] == '-' {
		sign = -1
		i++
	}

	var dec float64
	for i < len(bs) && bs[i] != '.' && bs[i] != ' ' {
		dec = dec*10 + float64(bs[i]-'0')
		i++
	}
	*val = dec

	if i == len(bs) || bs[i] == ' ' {
		//no fraction
		return i
	}
	i++
	var frac float64
	base := 1.0
	for i < len(bs) && bs[i] != ' ' {
		frac = frac*10 + float64(bs[i]-'0')
		base *= 10
		i++
	}
	*val += frac / base
	return i * sign
}

package main

import (
	"bufio"
	"bytes"
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		a, b, k := readThreeNums(reader)
		queries := make([][]int, k)
		for i := 0; i < k; i++ {
			queries[i] = readNNums(reader, 2)
		}
		res := solve(a, b, queries)

		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}
func solve(a, b int, queries [][]int) []int {
	ab := a * b
	pref := make([]int, ab+1)

	for i := 1; i <= ab; i++ {
		pref[i] = pref[i-1]
		if (i%a)%b != (i%b)%a {
			pref[i]++
		}
	}

	get := func(l int) int {
		cnt, rem := l/ab, l%ab
		return pref[ab]*cnt + pref[rem]
	}

	calc := func(l int, r int) int {
		return get(r) - get(l-1)
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		l, r := cur[0], cur[1]
		ans[i] = calc(l, r)
	}

	return ans
}

func lcm(a, b int) int {
	g := gcd(a, b)
	return a / g * b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

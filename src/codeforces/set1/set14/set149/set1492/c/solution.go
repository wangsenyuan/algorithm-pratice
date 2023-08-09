package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	s := readString(reader)[:n]
	t := readString(reader)[:m]

	res := solve(s, t)

	fmt.Println(res)
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

func solve(s, t string) int {
	//n := len(s)
	//m := len(t)
	// s[p1] = t[1]
	// s[p2] = t[2]
	// ...
	// 感觉从两头开始假设t[i]能够匹配的最短的前缀s[:p[i]]
	// 且 t[i+1]能够匹配最短的后缀s[x[i+1]:]
	pref := getMatchingPoints([]byte(t), []byte(s))
	suf := getMatchingPoints(reverse([]byte(t)), reverse([]byte(s)))
	n := len(s)
	m := len(t)

	suf = reverse(suf)

	for i := 0; i < m; i++ {
		suf[i] = n - 1 - suf[i]
	}

	var ans int
	for i := 0; i+1 < m; i++ {
		a := pref[i]
		b := suf[i+1]
		ans = max(ans, b-a)
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func getMatchingPoints(a []byte, b []byte) []int {
	n := len(a)
	m := len(b)
	res := make([]int, n)

	for i, j := 0, 0; i < n; i++ {
		for j < m && b[j] != a[i] {
			j++
		}
		// b[j] = a[i]
		res[i] = j
		j++
	}
	return res
}

type Item interface {
	int | byte
}

func reverse[T Item](arr []T) []T {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

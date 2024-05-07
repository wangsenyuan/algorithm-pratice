package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	s := readString(reader)[:n]
	res := solve(s, k)

	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

func solve(s string, k int) string {
	var set int
	n := len(s)
	for i := 0; i < n; i++ {
		x := int(s[i] - 'a')
		set |= 1 << x
	}

	getFirst := func() int {
		for i := 0; i < 26; i++ {
			if (set>>i)&1 == 1 {
				return i
			}
		}
		return 26
	}

	if k > n {
		x := getFirst()

		buf := make([]byte, k)
		copy(buf, s)
		for i := n; i < k; i++ {
			buf[i] = byte(x + 'a')
		}
		return string(buf)
	}
	// k <= n
	res := []byte(s[:k])
	for i := k - 1; i >= 0; i-- {
		x := int(s[i] - 'a')
		y := x + 1
		for y < 26 && (set>>y)&1 == 0 {
			y++
		}
		if y < 26 {
			res[i] = byte(y + 'a')
			z := getFirst()
			for j := i + 1; j < k; j++ {
				res[j] = byte(z + 'a')
			}
			break
		}
	}
	return string(res)
}

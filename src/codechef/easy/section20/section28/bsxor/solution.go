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
		n, k := readTwoNums(reader)
		s, _ := reader.ReadString('\n')
		res := solve(n, k, s)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}
	fmt.Print(buf.String())
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
		if s[i] == '\n' {
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

func solve(n int, k int, s string) bool {
	if k == 1 {
		return true
	}
	// xor s[l...r] = 0 or 1
	// let assume p0, p1, p2, ... p(k-1)
	// where p0 > 0 and p(k-1) == n
	// xor[0..p0] = 1
	// xor[p0..p1] = 1
	// xor[p1..p2] = 1
	// ...
	// xor[p(k-2)...p(k-1)] = 1
	// xor them will be 0 if k is even, else 1
	// so if xor of s = 1, and k is even, then no answer
	// if xor of s = 0,
	//         when k is even, partition could be 0 or 1
	//         when k is odd, partition could only be 0

	check := func(x int) bool {
		var cnt int
		var xor int
		var i int
		for ; i < n; i++ {
			xor ^= int(s[i] - '0')
			if xor == x {
				cnt++
				xor = 0
			}
			if cnt+1 == k {
				i++
				break
			}
		}
		if i == n {
			return false
		}
		// left last partition to check
		for i < n {
			xor ^= int(s[i] - '0')
			i++
		}
		return xor == x
	}

	var x int
	for i := 0; i < n; i++ {
		x ^= int(s[i] - '0')
	}

	if x == 1 {
		if k%2 == 0 {
			return false
		}

		return check(1)
	}
	// x == 0
	if k%2 == 1 {
		// only 0
		return check(0)
	}

	return check(0) || check(1)
}

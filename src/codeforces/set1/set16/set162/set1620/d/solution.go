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
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func readLenNums(r *bufio.Reader) []int {
	s, _ := r.ReadBytes('\n')
	var n int
	pos := readInt(s, 0, &n) + 1
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		pos = readInt(s, pos, &arr[i]) + 1
	}
	return arr
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(A []int) int {
	// any的意思是所有，也就说说，如果有rem = 1的， 那至少要带一个1，2至少要带一个2，或者在有1的时候，再带一个1
	// 除了要知道最小值，还需要知道最大值

	var res int
	for _, n := range A {
		res = max(res, n)
	}

	check := func(a, b, c int) bool {
		for _, n := range A {
			if !checkNum(a, b, c, n) {
				return false
			}
		}
		return true
	}

	ans := (res+2)/3 + 1
	// 是否可以有res个coins的结果呢？
	// a 个1
	for a := 0; a <= 3; a++ {
		// b 个2
		for b := 0; b <= 3; b++ {
			c := (res - a - 2*b) / 3
			if c < 0 {
				continue
			}
			if check(a, b, c) {
				ans = min(ans, a+b+c)
			}
		}
	}

	return ans
}

func checkNum(a, b, c int, num int) bool {
	// 优先使用3
	// let x, y, z => x + 2 * y + 3 * z = num
	// x <= a, y <= b, z <= c
	for x := 0; x <= a; x++ {
		for y := 0; y <= b; y++ {
			tmp := num - x - 2*y
			if tmp >= 0 && tmp%3 == 0 && tmp/3 <= c {
				return true
			}
		}
	}
	return false
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

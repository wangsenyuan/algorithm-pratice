package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	x := readNInt64s(reader, n)
	y := readNInt64s(reader, n)
	res := solve(x, y)
	fmt.Println(res)
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

const X = 15000010

func solve(x, y []int64) int {
	prime := make([]bool, X)
	prime_seg := make([]bool, X)

	for i := 0; i < X; i++ {
		prime_seg[i] = true
		prime[i] = true
	}

	L := min(x...) + min(y...)
	R := max(x...) + max(y...)

	for i := 2; i < X; i++ {
		if !prime[i] {
			continue
		}
		for j := 2 * i; j < X; j += i {
			prime[j] = false
		}
		lo := int64(i) * max(2, (L+int64(i)-1)/int64(i))
		for lo <= R {
			prime_seg[int(lo-L)] = false
			lo += int64(i)
		}
	}

	var ans int

	for i := 0; i < len(x); i++ {
		j := (i + 1) % len(x)
		dx := x[j] - x[i]
		dy := y[j] - y[i]
		g := gcd(abs(dx), abs(dy))
		a := dx / g
		b := dy / g
		cx, cy := x[i], y[i]
		for cx != x[j] || cy != y[j] {
			if prime_seg[int(cx+cy-L)] {
				ans++
			}
			cx += a
			cy += b
		}
	}

	return ans
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func max(arr ...int64) int64 {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
func min(arr ...int64) int64 {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res > arr[i] {
			res = arr[i]
		}
	}
	return res
}

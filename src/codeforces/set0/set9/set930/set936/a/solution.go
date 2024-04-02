package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	k, d, x := readThreeNums(reader)

	res := solve(k, d, x)

	fmt.Printf("%.9f\n", res)
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

func solve(k int, d int, t int) float64 {
	x := (k + d - 1) / d * d
	y := x - k

	check := func(expect float64) bool {
		// 在这么长的时间内，
		// 到底有多少时间是on的？
		// 有多少时间是off的？
		// 如果 k % d == 0, 那么所有的时间都是on的（off的瞬间被点燃了）
		// 假设在时刻x, julia 检查的时候，发现是off的，
		// 那么有 x % d = 0, 在这之前off的时间 = x - k
		// x = (k / d + 1) * d
		if k%d == 0 {
			return expect >= float64(t)
		}
		// 一个周期是这么长的
		w := int(expect)
		a, b := w/x, w%x
		on := float64(a * k)
		off := float64(a * y)
		f := float64(b) + expect - float64(w)
		if f <= float64(k) {
			on += f
		} else {
			on += float64(k)
			f -= float64(k)
			off += f
		}
		return on+off/2 >= float64(t)
	}

	var l, r float64 = 0, 1 << 62

	for i := 0; i < 100; i++ {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	return (l + r) / 2
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

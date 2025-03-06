package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	if len(res) == 0 {
		fmt.Println("-1")
	} else {
		fmt.Println(res[0], res[1])
	}
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

func process(reader *bufio.Reader) []int {
	x, y, w := readThreeNums(reader)
	return solve(x, y, w)
}

func solve(mx int, my int, w int) []int {
	freq := make(map[pair]int)

	// 但是如何找到最小值是个问题
	// 对于给定的x, 要找到最小的y， 满足这个数量 a * b = rev(a) * rev(b) >= w
	// 感觉可以用双指针？

	for y := 1; y <= my; y++ {
		ry := rev(y)
		g := gcd(ry, y)
		freq[pair{ry / g, y / g}]++
	}
	var cnt int
	x := 1

	freq1 := make(map[pair]int)
	for x <= mx {
		rx := rev(x)
		g := gcd(x, rx)
		key := pair{x / g, rx / g}
		if cnt+freq[key] >= w {
			break
		}
		cnt += freq[key]
		freq1[key]++
		x++
	}
	if x > mx {
		return nil
	}
	res := []int{x, my}

	for y := my; x <= mx; x++ {
		rx := rev(x)
		g := gcd(x, rx)
		cnt += freq[pair{x / g, rx / g}]
		freq1[pair{x / g, rx / g}]++
		// cnt >= w 成立
		for y > 0 {
			ry := rev(y)
			g := gcd(ry, y)
			key := pair{ry / g, y / g}
			if cnt-freq1[key] < w {
				break
			}
			freq[key]--
			cnt -= freq1[key]
			y--
		}
		// cnt >= w 始终是保持的
		if x*y < res[0]*res[1] {
			res[0] = x
			res[1] = y
		}
	}

	return res
}

func rev(num int) int {
	var res int
	for num > 0 {
		x := num % 10
		res = res*10 + x
		num /= 10
	}
	return res
}

type pair struct {
	first  int
	second int
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line := readNNums(reader, 5)
	P := line[0:2]
	V := line[2:4]
	k := line[4]
	res := solve(P, V, k)
	fmt.Printf("%s\n", res)
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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
		if s[i] == '\n' || s[i] == '\r' {
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

var lucky []int64

const MAX_X = 1e9

func init() {
	lucky = append(lucky, 0)
	for l := 1; l <= 9; l++ {
		for mask := 0; mask < 1<<l; mask++ {
			var sum int64
			for i := l - 1; i >= 0; i-- {
				sum = sum * 10
				if (mask>>i)&1 == 0 {
					sum += 4
				} else {
					sum += 7
				}
			}
			lucky = append(lucky, sum)
		}
	}

	lucky = append(lucky, MAX_X)
}

func solve(P []int, V []int, k int) string {
	var tot = int64(P[1]-P[0]+1) * int64(V[1]-V[0]+1)
	var ways int64
	pr := int64(P[1])
	pl := int64(P[0])
	vr := int64(V[1])
	vl := int64(V[0])
	// len(luck_nums) <= 1024
	for l, r := 1, k; r+1 < len(lucky); l, r = l+1, r+1 {
		//j := i + k - 1
		//检查是否可以和区间重叠
		//其中一个人选择luck[i] （包括）之前的数字
		//另外一个人选择 luck[j]之后的数字
		la := lucky[l-1] + 1
		ra := lucky[l]
		lb := lucky[r]
		rb := lucky[r+1] - 1

		//var ways int64
		ways += max(0, min(ra, pr)-max(la, pl)+1) * max(0, min(rb, vr)-max(lb, vl)+1)
		ways += max(0, min(ra, vr)-max(la, vl)+1) * max(0, min(rb, pr)-max(lb, pl)+1)
		if k == 1 && pl <= ra && ra <= pr && vl <= ra && ra <= vr {
			ways--
		}
		//ans += float64(ways) / float64(tot)
	}

	ans := big.NewFloat(0.0)
	ans = ans.Quo(big.NewFloat(float64(ways)), big.NewFloat(float64(tot)))

	return ans.String()
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type Num interface {
	int | int64
}

func max[T Num](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func min[T Num](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

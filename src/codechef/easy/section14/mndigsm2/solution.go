package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		var n, r uint64
		s, _ := reader.ReadBytes('\n')
		pos := readUint64(s, 0, &n)
		readUint64(s, pos+1, &r)
		res := solve(int64(n), int64(r))
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(n, r int64) int64 {

	if n <= r {
		return n
	}

	var ans int64 = 2
	best := digitSum(n, 2)

	// a * i ** 2 + b * i + c
	for i := int64(2); i*i*i <= n; i++ {
		if i > r {
			break
		}
		tmp := digitSum(n, i)
		if tmp < best {
			best = tmp
			ans = i
		}
	}
	// best = a + b + c <= lg(num)
	var y = best
	for a := int64(0); a <= y; a++ {
		for b := int64(0); b <= y; b++ {
			for c := int64(0); c <= y; c++ {
				if a == 0 {
					if b > 0 {
						B := (n - c) / b
						if B >= 2 && B <= r {
							tmp := digitSum(n, B)
							if tmp < best {
								best = tmp
								ans = B
							}
						}
					}
					continue
				}
				// a * x ** 2 + b * x + c = n
				det := b*b + 4*a*(n-c)
				// det > 0
				sqt := int64(math.Sqrt(float64(det)))
				if sqt*sqt < det {
					sqt++
				}
				B := (sqt - b) / (2 * a)
				if B >= 2 && B <= r {
					tmp := digitSum(n, B)
					if tmp < best {
						best = tmp
						ans = B
					}
				}
			}
		}
	}

	return ans
}

func digitSum(num int64, B int64) int64 {

	var res int64

	for num > 0 {
		res += num % B
		num /= B
	}
	return res
}

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
		num := readNum(reader)
		res := solve(num)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

const M = 99999 + 1

var primes []bool
var pw [10]int

func init() {
	primes = make([]bool, M)
	for i := 2; i < M; i++ {
		primes[i] = true
	}
	for i := 2; i < M; i++ {
		if primes[i] {
			for j := i * 2; j < M; j += i {
				primes[j] = false
			}
		}
	}
	pw[0] = 1
	for i := 1; i < 10; i++ {
		pw[i] = 10 * pw[i-1]
	}
}

func solve(first int) int {
	arr := numToArray(first)
	n := len(arr)
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}
	copy(grid[0], arr)

	var dfs func(r int) int

	dfs = func(r int) int {
		if r == n {
			return 1
		}
		var num int
		for i := 0; i < r; i++ {
			num = num*10 + grid[i][r]
		}
		num *= pw[n-r]
		var ans int
		for i := 0; i < pw[n-r-1]; i++ {
			var cnt int
			for j := 0; j < 10; j++ {
				if primes[num+j*pw[n-r-1]+i] {
					cnt++
				}
			}
			if cnt != 0 {
				tmp := i
				for j := n - 1; j > r; j-- {
					grid[r][j] = tmp % 10
					tmp /= 10
				}
				ans += cnt * dfs(r+1)
			}
		}
		return ans
	}

	return dfs(1)
}

func numToArray(num int) []int {
	var res []int
	for num > 0 {
		res = append(res, num%10)
		num /= 10
	}
	reverse(res)
	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

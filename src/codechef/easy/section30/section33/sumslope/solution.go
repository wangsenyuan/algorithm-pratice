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
		var A, B uint64
		s, _ := reader.ReadBytes('\n')
		pos := readUint64(s, 0, &A)
		readUint64(s, pos+1, &B)
		res := solve(int64(A), int64(B))
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(A, B int64) int64 {
	return countSlope(B) - countSlope(A-1)
}

const H = 19

var PW []int64

const X = 1000

var C []int
var A []bool

func init() {
	// F[x][i] 表示以x为head，且长度为i的数字，且下个数字y < x的计数
	// x (x + 1) x
	PW = make([]int64, H)
	PW[0] = 1
	for i := 1; i < H; i++ {
		PW[i] = 10 * PW[i-1]
	}

	C = make([]int, X)
	A = make([]bool, X)

	for x := 1; x < X; x++ {
		a := x / 100
		b := (x % 100) / 10
		c := x % 10
		A[x] = slop(a, b, c)
		if A[x] {
			C[x]++
		}
		C[x] += C[x-1]
	}
}

func countSlope(n int64) int64 {
	if n <= 100 {
		return 0
	}
	if n < X {
		return int64(C[int(n)] - C[100])
	}

	B := make([]int, H)
	suf := make([]int64, H)
	var ptr int

	for n > 0 {
		B[ptr] = int(n % 10)
		ptr++

		suf[ptr] = suf[ptr-1] + int64(B[ptr-1]*B[ptr-1])

		n /= 10
	}

	var ans int64
	var tail int64

	for i := ptr - 1; i >= 2; i-- {
		k := i
		if tail > 0 {
			ans += (tail - 1) * 570 * PW[k-2]
			ans += int64(570-C[100]) * PW[k-2]
		}
		v := B[k]*100 + B[k-1]*10 + B[k-2]

		if v > 0 {
			if tail > 0 {
				ans += int64(C[v-1]) * PW[k-2]
			} else {
				if C[v-1] > C[100] {
					ans += int64(C[v-1]-C[100]) * PW[k-2]
				}
			}
		}
		if A[v] {
			ans += suf[k-2] + 1
		}
		tail = tail*10 + int64(B[i])
	}

	return ans
}

func countSlope1(num int64) int64 {
	if num <= 100 {
		return 0
	}
	tmp := num
	var l int
	arr := make([]int, H)
	for num > 0 {
		arr[l] = int(num % 10)
		num /= 10
		l++
	}
	reverse(arr[:l])

	dp := make([][][]int64, l)
	for i := 0; i < l; i++ {
		dp[i] = make([][]int64, 100)
		for j := 0; j < 100; j++ {
			dp[i][j] = make([]int64, 4)
		}
	}

	for x := 0; x < 100; x++ {
		d1 := x / 10
		d2 := x % 10
		if d1 > arr[0] || d1 == arr[0] && d2 > arr[1] {
			break
		}
		var eq int
		if d1 == arr[0] && d2 == arr[1] {
			eq = 2
		}
		var leading_zero int
		if d1 == 0 {
			leading_zero = 1
		}
		dp[1][x][eq+leading_zero]++
	}

	var res int64

	for i := 2; i < l; i++ {
		// try to put d3 at position i
		for x := 0; x < 100; x++ {
			d1 := x / 10
			d2 := x % 10
			for u := 0; u < 4; u++ {
				if dp[i-1][x][u] == 0 {
					continue
				}
				eq := u / 2
				leading_zero := u % 2
				var next_leading_zero = 1
				if leading_zero == 0 || d2 > 0 {
					next_leading_zero = 0
				}
				for d3 := 0; d3 < 10; d3++ {
					if eq == 0 {
						if leading_zero == 0 && slop(d1, d2, d3) {
							res += PW[l-i-1] * dp[i-1][x][u]
						}
						dp[i][d2*10+d3][next_leading_zero] += dp[i-1][x][u]
					} else {
						// eq = 1, zr must be 0
						if d3 < arr[i] {
							if leading_zero == 0 && slop(d1, d2, d3) {
								res += PW[l-i-1] * dp[i-1][x][u]
							}
							// not eq any more
							dp[i][d2*10+d3][next_leading_zero] += dp[i-1][x][u]
						} else if d3 == arr[i] {
							// d3 == arr[i], still eq
							if slop(d1, d2, d3) {
								res += (tmp%PW[l-i-1] + 1) * dp[i-1][x][u]
							}
							dp[i][d2*10+d3][2+next_leading_zero] += dp[i-1][x][u]
						}
					}
				}
			}
		}
	}

	return res
}

func slop(a, b, c int) bool {
	return a < b && b > c || a > b && b < c
}

func count(num int64) int64 {
	var arr []int
	for num > 0 {
		arr = append(arr, int(num%10))
		num /= 10
	}
	var cnt int64
	for i := 1; i+1 < len(arr); i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			cnt++
		} else if arr[i] < arr[i-1] && arr[i] < arr[i+1] {
			cnt++
		}
	}
	return cnt
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

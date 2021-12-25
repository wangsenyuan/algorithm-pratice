package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		s, _ := reader.ReadString('\n')
		s = s[:len(s)-1]
		fmt.Println(solve(s))
	}
}

const MAX_N = 100005

var Z [][]int

func init() {
	Z = make([][]int, 2)
	for i := 0; i < 2; i++ {
		Z[i] = make([]int, MAX_N)
	}
}

func reset(n int) {
	for i := 0; i < 2; i++ {
		for j := 0; j <= n; j++ {
			Z[i][j] = 0
		}
	}
}

func solve(s string) int {
	n := len(s)
	reset(n)

	zFunction(s, Z[0])

	zFunction(reverse(s), Z[1])

	var ans int

	for i := 1; i < n; i += 2 {
		j := n - i - 2
		if j < 0 {
			break
		}
		if Z[0][i/2+1] >= (i+1)/2 && Z[1][j/2+1] >= (j+1)/2 {
			ans++
		}
	}

	return ans
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func zFunction(s string, z []int) {
	n := len(s)
	z[0] = 0
	var l, r int
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[i+z[i]] == s[z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			r = i + z[i] - 1
			l = i
		}
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// var B []int64
// var M []int64
// var H [][]int64

// const MAX_N = 100007

// var BB [][]int64

// func init() {
// 	B = make([]int64, 3)
// 	M = make([]int64, 3)
// 	B[0] = 27
// 	M[0] = 1000000007
// 	B[1] = 28
// 	M[1] = 1000000009
// 	B[2] = 29
// 	M[2] = 10000000007

// 	H = make([][]int64, 3)
// 	BB = make([][]int64, 3)

// 	for i := 0; i < 3; i++ {
// 		BB[i] = make([]int64, MAX_N)
// 		BB[i][0] = 1
// 		H[i] = make([]int64, MAX_N)
// 		H[i][0] = 0
// 	}

// 	for j := 1; j < MAX_N; j++ {
// 		for i := 0; i < 3; i++ {
// 			BB[i][j] = (BB[i][j-1] * B[i]) % M[i]
// 		}
// 	}
// }

// func solve(s string) int {
// 	n := len(s)

// 	for i := 0; i < n; i++ {
// 		for j := 0; j < 3; j++ {
// 			H[j][i+1] = (H[j][i]*B[j] + int64(s[i]-'a'+1)) % M[j]
// 		}
// 	}

// 	var res int

// 	for i := 1; i < n; i++ {
// 		// i is the length of T1
// 		ii := i + i

// 		if n == ii || (n-ii)%2 != 0 {
// 			continue
// 		}

// 		var can = true

// 		for j := 0; j < 3 && can; j++ {
// 			h1 := H[j][i]
// 			h2 := H[j][ii]
// 			h2 -= (h1 * BB[j][i]) % M[j]
// 			if h2 < 0 {
// 				h2 += M[j]
// 			}
// 			if h1 != h2 {
// 				can = false
// 				break
// 			}
// 			l := (n - ii) / 2
// 			k := ii + l
// 			h3 := H[j][k]
// 			h3 -= (H[j][ii] * BB[j][l]) % M[j]
// 			if h3 < 0 {
// 				h3 += M[j]
// 			}
// 			h4 := H[j][n]
// 			h4 -= (H[j][k] * BB[j][l]) % M[j]
// 			if h4 < 0 {
// 				h4 += M[j]
// 			}
// 			can = h3 == h4
// 		}

// 		if can {
// 			res++
// 		}
// 	}

// 	return res
// }

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
		res := solve(k, n)
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func solve(k int, n int) []int {
	// dp[x][sz][xor]
	// p[x][sz][xor]
	dp := make([][]map[int]bool, n+1)
	fp := make([][]map[int]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]map[int]bool, k+1)
		fp[i] = make([]map[int]bool, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make(map[int]bool)
			fp[i][j] = make(map[int]bool)
		}
	}
	dp[0][0][0] = true

	for x := 0; x < n; x++ {
		for sz := 0; sz <= x && sz <= k; sz++ {
			// xor would be
			y := highestOneBit(x) * 2
			for xor := 0; xor <= y; xor++ {
				if dp[x][sz][xor] {
					if !dp[x+1][sz][xor] {
						dp[x+1][sz][xor] = true
						fp[x+1][sz][xor] = false
					}
					if sz+1 <= k && !dp[x+1][sz+1][xor^(x+1)] {
						dp[x+1][sz+1][xor^(x+1)] = true
						fp[x+1][sz+1][xor^(x+1)] = true
					}
				}
			}
		}
	}

	top := highestOneBit(n) * 2
	var best int
	for xor := 0; xor <= top; xor++ {
		if dp[n][k][xor] {
			best = xor
		}
	}

	var arr []int

	for pos, sz := n, k; pos > 0; pos-- {
		if fp[pos][sz][best] {
			arr = append(arr, pos)
			best ^= pos
			sz--
		}
	}

	reverse(arr)

	return arr
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func solveFull(k int, n int) []int {
	if k == 1 {
		return []int{n}
	}
	if k == n {
		res := make([]int, n)
		for i := 0; i < n; i++ {
			res[i] = i + 1
		}
		return res
	}
	B := highestOneBit(n)

	if k == 2 {
		// find some b, 2 ^^ b <= n
		return []int{B - 1, B}
	}

	if k+1 == n {
		var xor int
		for i := 1; i <= n; i++ {
			xor ^= i
		}

		// find some x, that makes xor ^ x maximized
		pick := 1
		for i := 1; i < n; i++ {
			if xor^pick < xor^(i+1) {
				pick = i + 1
			}
		}
		res := make([]int, 0, n-1)
		for i := 0; i < n; i++ {
			if pick != (i + 1) {
				res = append(res, i+1)
			}
		}
		return res
	}
	maxAns := (highestOneBit(n) << 1) - 1

	if k+2 == n {
		res := make([]int, 0, k)

		var xor int
		for i := 1; i <= n; i++ {
			xor ^= i
		}

		xor ^= maxAns

		if bitCount(xor) >= 2 {
			a := highestOneBit(xor)
			b := xor ^ a
			for i := 1; i <= n; i++ {
				if i != a && i != b {
					res = append(res, i)
				}
			}
			return res
		}
		if bitCount(xor) == 1 {
			var a, b int
			if xor == 1 {
				a, b = 2, 3
			} else {
				a = xor ^ 1
				b = 1
			}

			for i := 1; i <= n; i++ {
				if i != a && i != b {
					res = append(res, i)
				}
			}
			return res
		} else {
			a := highestOneBit(n)
			b := a + 1
			for i := 1; i <= n; i++ {
				if i != a && i != b {
					res = append(res, i)
				}
			}
			return res
		}
	}
	var xor int
	inc := make([]bool, n+1)
	for i := 1; i <= k; i++ {
		xor ^= i
		inc[i] = true
	}
	if k < B-1 {
		btc := bitCount(xor)
		if btc >= 2 {
			inc[B] = true
			inc[B-1] = true
			xor = xor ^ (B - 1) ^ B ^ maxAns
			a := highestOneBit(xor)
			b := xor ^ a
			inc[a] = false
			inc[b] = false
		} else if btc == 1 {
			if xor == 1 {
				inc[B] = true
				inc[B-1] = true
				inc[2] = false
				inc[3] = false
			} else {
				inc[B-2] = true
				inc[B] = true
				inc[xor] = false
				inc[1] = false
			}
		} else {
			inc[B-2] = true
			inc[B] = true
			inc[2] = false
			inc[3] = false
		}
	} else if k == B-1 {
		inc[B] = true
		inc[B-1] = false
	} else {
		if xor&B == 0 {
			inc[k+1] = true
			xor ^= (k + 1)
			inc[xor^maxAns] = false
		} else {
			if xor^(k+1)&^(k+2) == maxAns {
				inc[k+1] = true
				inc[k+3] = true
				xor = xor ^ (k + 1) ^ (k + 3)
			} else {
				inc[k+1] = true
				inc[k+2] = true
				xor = xor ^ (k + 1) ^ (k + 2)
			}
			xor ^= maxAns
			if bitCount(xor) >= 2 {
				a := highestOneBit(xor)
				b := xor ^ a
				inc[a] = false
				inc[b] = false
			} else {
				if xor == 1 {
					inc[2] = false
					inc[3] = false
				} else {
					inc[xor+1] = false
					inc[1] = false
				}
			}
		}
	}

	res := make([]int, 0, k)
	for i := 1; i <= n; i++ {
		if inc[i] {
			res = append(res, i)
		}
	}
	return res
}

func bitCount(num int) int {
	if num == 0 {
		return 0
	}
	return 1 + bitCount(num&(num-1))
}

func highestOneBit(i int) int {
	return i & (-2147483648 >> uint(numberOfLeadingZeros(i)))
}

func numberOfLeadingZeros(i int) int {
	if i <= 0 {
		if i == 0 {
			return 32
		}
		return 0
	}
	n := 31
	if i >= 65536 {
		n -= 16
		i >>= 16
	}

	if i >= 256 {
		n -= 8
		i >>= 8
	}

	if i >= 16 {
		n -= 4
		i >>= 4
	}

	if i >= 4 {
		n -= 2
		i >>= 2
	}

	return n - (i >> 1)
}

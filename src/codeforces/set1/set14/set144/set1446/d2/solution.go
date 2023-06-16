package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)

	fmt.Println(solve(A))
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
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
	// A[i] <= min(n, 100)
	// 假设到r为止的freq[x] 表示x的频率
	// 如果知道y，freq[x] - freq[y] 的最小值
	n := len(A)
	m := A[0]
	for i := 0; i < n; i++ {
		m = max(m, A[i])
	}

	m++

	freq := make([]int, m)

	for _, num := range A {
		freq[num]++
	}
	d := 1
	for i := 1; i < m; i++ {
		if freq[d] < freq[i] {
			d = i
		}
	}

	for j := 1; j < m; j++ {
		if j != d && freq[j] == freq[d] {
			return n
		}
	}

	if freq[d] == n {
		return 0
	}

	sq := int(math.Sqrt(float64(n)))

	pos := make([]int, 2*n+1)

	var best int

	for v := 1; v < m; v++ {
		if v == d || freq[v] <= sq {
			continue
		}
		for j := 0; j < len(pos); j++ {
			pos[j] = n
		}
		pos[n] = -1
		sum := n
		for j := 0; j < n; j++ {
			if A[j] == d {
				sum--
			} else if A[j] == v {
				sum++
			}

			best = max(best, j-pos[sum])
			pos[sum] = min(pos[sum], j)
		}
	}

	for v := 1; v <= sq; v++ {
		for i := 0; i < m; i++ {
			freq[i] = 0
		}
		for j := 0; j < len(pos); j++ {
			pos[j] = 0
		}
		for l, r := 0, 0; r < n; r++ {
			pos[freq[A[r]]]--
			freq[A[r]]++
			pos[freq[A[r]]]++
			for freq[A[r]] > v {
				pos[freq[A[l]]]--
				freq[A[l]]--
				pos[freq[A[l]]]++
				l++
			}
			if pos[v] >= 2 {
				best = max(best, r-l+1)
			}
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	return a + b - max(a, b)
}

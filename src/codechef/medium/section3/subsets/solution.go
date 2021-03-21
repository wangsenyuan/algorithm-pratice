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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(n, A)
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func solve(n int, A []int) int {
	mat := make([][]bool, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]bool, n)
	}

	firstPart := make([]int, n)
	half := n / 2

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if adjacent(A[i], A[j]) {
				if j < n {
					firstPart[i] |= 1 << uint(j)
				}
				if i < n {
					firstPart[j] |= 1 << uint(i)
				}
				mat[i][j] = true
				mat[j][i] = true
			}
		}
	}

	clique := make([]bool, 1<<uint(half))
	clique[0] = true

	for mask := 1; mask < 1<<uint(half); mask++ {
		v := ctz(mask)
		mask2 := mask ^ (1 << uint(v))
		if !clique[mask2] {
			clique[mask] = false
			continue
		}
		clique[mask] = true
		for i := 0; i < half; i++ {
			if mask2&(1<<uint(i)) > 0 && !mat[v][i] {
				clique[mask] = false
				break
			}
		}
	}

	dp := make([]int, 1<<uint(n-half))
	dp[0] = 1

	for mask := 1; mask < 1<<uint(n-half); mask++ {
		v := ctz(mask)
		mask2 := mask ^ (1 << uint(v))
		dp[mask] = dp[mask2]
		var mask3 int
		for i := 0; i < n-half; i++ {
			if mat[v+half][i+half] && mask&(1<<uint(i)) > 0 {
				mask3 |= 1 << uint(i)
			}
		}
		dp[mask] += dp[mask3]
	}

	var ans int

	for mask := 0; mask < 1<<uint(half); mask++ {
		if !clique[mask] {
			continue
		}
		var mask2 int
		for i := half; i < n; i++ {
			if firstPart[i]&mask == mask {
				mask2 |= 1 << uint(i-half)
			}
		}
		ans += dp[mask2]
	}
	return ans
}

func ctz(num int) int {
	var i int
	for num&1 == 0 {
		i++
		num >>= 1
	}
	return i
}

func adjacent(a, b int) bool {
	factors := make(map[int]int)

	factorize(a, factors)
	factorize(b, factors)

	for _, v := range factors {
		if v >= 3 {
			return false
		}
	}
	return true
}

func factorize(num int, fs map[int]int) {
	for x := 2; x*x <= num; x++ {
		for num%x == 0 {
			fs[x]++
			if fs[x] >= 3 {
				return
			}
			num /= x
		}
	}
	if num > 1 {
		fs[num]++
	}
}

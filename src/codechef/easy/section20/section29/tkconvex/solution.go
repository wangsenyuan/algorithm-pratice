package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	n, k := readTwoNums(reader)
	A := readNNums(reader, n)
	res := solve(n, k, A)
	if len(res) > 0 {
		buf.WriteString(fmt.Sprintf("Yes\n"))
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]+1))
		}
		buf.WriteByte('\n')
	} else {
		buf.WriteString("No\n")
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, k int, A []int) []int {
	B := make([]Pair, n)

	for i := 0; i < n; i++ {
		B[i] = Pair{A[i], i}
	}

	sort.Sort(Pairs(B))

	var sum int64
	first := -1

	res := make([]int, 2*k)

	arr := make([]int, 2*k)

	for i := 0; i <= n; i++ {
		if i >= k {
			// sum is s[i-k:i]
			if 2*int64(B[i-1].first) < sum {
				// a good choice
				if first < 0 {
					first = i - 1
					for p := 0; p < k; p++ {
						res[p] = B[i-k+p].second
					}
				} else if i-2*k >= 0 {
					// already first >= 0
					if i-k > first {
						// two segments, one ending at first, another ending at first
						for p := 0; p < k; p++ {
							res[p+k] = B[i-k+p].second
						}
						return res
					}
					for p := 0; p < 2*k; p++ {
						arr[p] = B[i-2*k+p].first
					}
					if x := check(arr); x > 0 {
						// case two
						var p int
						for j := 0; j < 2*k; j++ {
							if (x>>uint(j))&1 == 1 {
								res[p] = B[i-2*k+j].second
								p++
							}
						}
						for j := 0; j < 2*k; j++ {
							if (x>>uint(j))&1 == 0 {
								res[p] = B[i-2*k+j].second
								p++
							}
						}
						return res
					}
				}
			}
			sum -= int64(B[i-k].first)
		}
		if i < n {
			sum += int64(B[i].first)
		}
	}

	return nil
}

type Pair struct {
	first  int
	second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first < this[j].first
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func check(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	// n <= 20
	k := n / 2

	x := (1 << uint(k)) - 1
	end := 1 << uint(n)

	for x < end {
		var first int64
		var second int64
		var a, b int64
		for i := 0; i < n; i++ {
			if (x>>uint(i))&1 == 1 {
				first += int64(arr[i])
				a = int64(arr[i])
			} else {
				second += int64(arr[i])
				b = int64(arr[i])
			}
		}
		if 2*a < first && 2*b < second {
			return x
		}
		x = nextNum(x)
	}
	return 0
}

func nextNum(x int) int {
	// right most set bit
	rightOne := x & -x

	// reset the pattern and set next higher bit
	// left part of x will be here
	nextHigherOneBit := x + rightOne

	// nextHigherOneBit is now part [D] of the above explanation.

	// isolate the pattern
	rightOnesPattern := x ^ nextHigherOneBit

	// right adjust pattern
	rightOnesPattern = (rightOnesPattern) / rightOne

	// correction factor
	rightOnesPattern >>= 2

	// rightOnesPattern is now part [A] of the above explanation.

	// integrate new pattern (Add [D] and [A])
	next := nextHigherOneBit | rightOnesPattern

	return next
}

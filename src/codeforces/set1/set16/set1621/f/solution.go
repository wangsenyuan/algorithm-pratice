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
	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		line := readNNums(reader, 4)
		s, _ := reader.ReadString('\n')
		res := solve(line[0], line[1], line[2], line[3], s)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func solve(n int, a, b, c int, s string) int64 {
	if n == 1 {
		return 0
	}
	fst1, lst1 := -1, -1

	for i := 0; i < n; i++ {
		if s[i] == '1' {
			if fst1 < 0 {
				fst1 = i
			}
			lst1 = i
		}
	}
	if fst1 < 0 {
		return int64(a)
	}

	var P int
	if s[0] == '0' {
		P++
	}

	if s[n-1] == '0' {
		P++
	}

	other0 := max(fst1-1, 0) + max(n-lst1-2, 0)

	var turn1 int

	for i := 0; i+1 < n; i++ {
		if s[i] == '1' && s[i] == s[i+1] {
			turn1++
		}
	}

	var blocks []int
	var single0 int

	for i := fst1; i < lst1; {
		j := i + 1
		for s[j] == '0' {
			j++
		}
		l := j - i - 1
		if l == 1 {
			single0++
		} else if l > 1 {
			blocks = append(blocks, l)
		}
		i = j
	}

	sort.Ints(blocks)
	reverse(blocks)

	y := process(1, int64(a), int64(b), int64(c), blocks[:], other0, single0, P, turn1)
	x := process(0, int64(a), int64(b), int64(c), blocks[:], other0, single0, P, turn1)

	return max2(x, y)
}

func process(turn int, a, b, c int64, blocks []int, other0 int, single0 int, P int, one int) (ans int64) {
	l := len(blocks)

	var cur int64
	for {
		if turn == 1 {
			if one == 0 {
				return
			}
			one--
			cur += b
			ans = max2(ans, cur)
		} else {
			if one == 0 {
				if other0 > 0 || l > 0 {
					ans = max2(ans, cur+a)
				}
				if single0 > 0 {
					single0--
					cur -= c
					one++
				}
			} else {
				if l > 0 {
					blocks[l-1]--
					if blocks[l-1] == 1 {
						l--
						single0++
					}
					cur += a
					ans = max2(ans, cur)
				} else if other0 > 0 {
					other0--
					cur += a
					ans = max2(ans, cur)
				} else if single0 > 0 {
					single0--
					cur -= c
					one++
				} else if P > 0 {
					P--
					cur -= c
				} else {
					return
				}
			}
		}
		turn ^= 1
	}
}

func max2(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

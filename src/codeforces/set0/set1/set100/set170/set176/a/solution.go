package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m, k := readThreeNums(reader)
	items := make([][][]int, n)
	for i := 0; i < n; i++ {
		readString(reader)
		items[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			items[i][j] = readNNums(reader, 3)
		}
	}

	res := solve(k, items)
	fmt.Println(res)
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

func solve(k int, goods [][][]int) int {
	n := len(goods)
	m := len(goods[0])
	var best int
	items := make([]Item, m)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for x := 0; x < m; x++ {
				items[x] = Item{x, goods[i][x][2], goods[j][x][1] - goods[i][x][0]}
			}
			sort.Slice(items, func(a, b int) bool {
				return items[a].profit > items[b].profit
			})
			var tmp, cnt int
			for x := 0; x < m && items[x].profit > 0 && cnt < k; x++ {
				if cnt+items[x].cnt <= k {
					tmp += items[x].cnt * items[x].profit
					cnt += items[x].cnt
				} else {
					tmp += (k - cnt) * items[x].profit
					cnt += k - cnt
				}
			}
			best = max(best, tmp)
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

type Item struct {
	id     int
	cnt    int
	profit int
}

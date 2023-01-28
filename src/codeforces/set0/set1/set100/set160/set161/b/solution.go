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
	n, k := readTwoNums(reader)
	items := make([][]int, n)
	for i := 0; i < n; i++ {
		items[i] = readNNums(reader, 2)
	}
	price, res := solve(k, items)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%.1f\n", price))
	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d", len(ans)))
		for i := 0; i < len(ans); i++ {
			buf.WriteString(fmt.Sprintf(" %d", ans[i]))
		}
		buf.WriteByte('\n')
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
func solve(k int, items [][]int) (float64, [][]int) {
	//n := len(items)
	var stools []Item
	var pencils []Item
	for i, it := range items {
		if it[1] == 1 {
			stools = append(stools, Item{it[0], i + 1})
		} else {
			pencils = append(pencils, Item{it[0], i + 1})
		}
	}
	sort.Slice(stools, func(i, j int) bool {
		return stools[i].price > stools[j].price
	})

	sort.Slice(pencils, func(i, j int) bool {
		return pencils[i].price < pencils[j].price
	})
	var pay float64
	carts := make([][]int, k)
	var u, v int
	for i := 0; i < k; i++ {
		carts[i] = make([]int, 0, 1)
		if u == len(stools) || i+1 == k {
			// all stools used up, or the last one
			if u < len(stools) {
				// i + 1 == k
				for u < len(stools) {
					carts[i] = append(carts[i], stools[u].id)
					pay += float64(stools[u].price)
					u++
				}
			}
			if v < len(pencils) {
				carts[i] = append(carts[i], pencils[v].id)
				pay += float64(pencils[v].price)
				v++
				for i+1 == k && v < len(pencils) {
					// last one
					carts[i] = append(carts[i], pencils[v].id)
					pay += float64(pencils[v].price)
					v++
				}
			}
			ok := false
			min_price := 1 << 30
			for x := 0; x < len(carts[i]); x++ {
				ok = ok || items[carts[i][x]-1][1] == 1
				if min_price > items[carts[i][x]-1][0] {
					min_price = items[carts[i][x]-1][0]
				}
			}
			if ok {
				pay -= float64(min_price) / 2
			}
		} else {
			carts[i] = append(carts[i], stools[u].id)
			pay += float64(stools[u].price) / 2
			u++
		}
	}
	return pay, carts
}

type Item struct {
	price int
	id    int
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

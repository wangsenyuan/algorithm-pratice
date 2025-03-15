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
	for range tc {
		k := readNum(reader)
		res := solve(k)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for _, cur := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
		}
	}
	buf.WriteTo(os.Stdout)
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

const inf = 1_000_000_000

func solve(k int) [][]int {
	var res [][]int

	x, y := -inf, -inf

	for k > 0 {
		// n * (n - 1) / 2 <= k
		n := sort.Search(k+2, func(i int) bool {
			return i*(i-1)/2 > k
		})
		n--
		for i := range n {
			res = append(res, []int{x + i, y})
		}
		x += n
		y += 2
		k -= n * (n - 1) / 2
	}

	return res
}

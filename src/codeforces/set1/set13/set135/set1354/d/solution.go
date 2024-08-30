package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, q := readTwoNums(reader)

	a := readNNums(reader, int(n))
	queries := readNNums(reader, int(q))

	res := solve(a, queries)

	fmt.Println(res)
}

func readInt(bytes []byte, from int, val *int32) int {
	i := from
	var sign int32 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int32
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int32(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int32) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int32, b int32) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int32, b int32, c int32) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int32 {
	res := make([]int32, n)
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

func solve(a []int32, queries []int32) int32 {

	count := func(x int32) int32 {
		var cnt int32
		for _, num := range a {
			if num <= x {
				cnt++
			}
		}
		for _, k := range queries {
			if k > 0 && k <= x {
				cnt++
			}
			if k < 0 && -k <= cnt {
				cnt--
			}
		}
		return cnt
	}

	if count(1<<30) == 0 {
		return 0
	}
	var l, r int32 = 0, 1 << 30
	for l < r {
		mid := (l + r) / 2
		if count(mid) > 0 {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

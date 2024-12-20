package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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


func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	queries := make([]int, m)
	for i := range m {
		queries[i] = readNum(reader)
	}
	return solve(a, queries)

}
func solve(a []int, queries []int) []int {
	// n := len(a)
	cnt := make([]int, 31)

	for _, num := range a {
		d := bits.Len(uint(num))
		cnt[d-1]++
	}

	find := func(num int) int {
		var res int
		var need int
		for i := 30; i >= 0; i-- {
			if (num>>i)&1 == 1 {
				need++
			}
			if cnt[i] >= need {
				res += need
				need = 0
			} else {
				res += cnt[i]
				need -= cnt[i]
			}
			need *= 2
		}
		if need > 0 {
			return -1
		}
		return res
	}

	ans := make([]int, len(queries))

	for i, num := range queries {
		ans[i] = find(num)
	}
	return ans
}

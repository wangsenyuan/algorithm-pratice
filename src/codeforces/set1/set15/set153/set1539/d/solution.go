package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	prods := make([][]int, n)
	for i := 0; i < n; i++ {
		prods[i] = readNNums(reader, 2)
	}
	res := solve(prods, n)
	fmt.Println(res)
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

func solve(products [][]int, n int) int {

	sort.Slice(products, func(i, j int) bool {
		return products[i][1] < products[j][1] || products[i][1] == products[j][1] && products[i][0] > products[j][0]
	})
	var cnt int
	var ans int
	for i, j := 0, n-1; i <= j; {
		for i <= j && cnt < products[i][1] {
			x := min(products[i][1]-cnt, products[j][0])
			if cnt < products[j][1] {
				ans += 2 * x
			} else {
				ans += x
			}
			cnt += x
			products[j][0] -= x
			if products[j][0] == 0 {
				j--
			}
		}
		if i <= j {
			// cnt >= b[i] 肯定成立
			ans += products[i][0]
			cnt += products[i][0]

			i++
		}
	}
	return ans
}

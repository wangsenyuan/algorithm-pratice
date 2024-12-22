package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)
	fmt.Println(res)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(arr ...int) int {
	res := 1
	for _, x := range arr {
		res = res * x % mod
	}
	return res
}

func solve(nums []int) int {
	pw := make([]int, 21)
	pw[0] = 1
	for i := 1; i < 21; i++ {
		pw[i] = mul(pw[i-1], 10)
	}

	cnt := make([]int, 11)
	n := len(nums)
	ds := make([][]int, n)

	for i := range n {
		ds[i] = toDigits(nums[i])
		cnt[len(ds[i])]++
	}

	var res int

	for _, cur := range ds {

		for d := 1; d <= 10; d++ {
			if cnt[d] == 0 {
				continue
			}
			var sum int
			var i int
			for i < len(cur) && i < d {
				sum = add(sum, mul(cur[i], pw[2*i+1]))
				i++
			}
			for i < len(cur) {
				sum = add(sum, mul(cur[i], pw[2*d+i-d]))
				i++
			}
			res = add(res, mul(sum, cnt[d]))

			sum = 0
			i = 0
			for i < len(cur) && i < d {
				sum = add(sum, mul(cur[i], pw[2*i]))
				i++
			}

			for i < len(cur) {
				sum = add(sum, mul(cur[i], pw[2*d+i-d]))
				i++
			}
			res = add(res, mul(sum, cnt[d]))
		}
	}

	return res
}

func toDigits(num int) []int {
	var arr []int
	for num > 0 {
		arr = append(arr, num%10)
		num /= 10
	}
	return arr
}

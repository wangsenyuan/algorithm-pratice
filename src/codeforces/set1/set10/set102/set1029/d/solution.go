package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)

	a := readNNums(reader, n)

	res := solve(a, k)

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

// 1000000000

func solve(a []int, k int) int {
	freq := make([]map[int]int, 11)

	base := make([]int, 11)
	base[0] = 1
	for i := 1; i < 11; i++ {
		freq[i] = make(map[int]int)
		base[i] = base[i-1] * 10
		base[i] %= k
	}

	for _, num := range a {
		for l := 1; l < 11; l++ {
			y := num % k * base[l] % k
			freq[l][y]++
		}
	}

	var res int

	//n := len(a)

	for _, num := range a {
		x := num % k
		ln := getLength(num)
		if x == 0 {
			res += freq[ln][0]
		} else {
			res += freq[ln][k-x]
		}

		if (x*base[ln]%k+x)%k == 0 {
			res--
		}
	}
	return res
}

func getLength(num int) int {
	var res int
	for num > 0 {
		res++
		num /= 10
	}
	return res
}

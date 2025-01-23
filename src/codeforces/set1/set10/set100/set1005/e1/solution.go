package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	p := readNNums(reader, n)
	return solve(m, p)
}

func solve(m int, p []int) int {
	n := len(p)
	var pm int
	m--
	for i := range n {
		p[i]--
		if p[i] == m {
			pm = i
		}
	}

	cnt := make([]int, 2*n)

	add := func(x int) {
		cnt[x+n]++
	}
	// 右边长度为0的情况
	add(0)
	var sum int
	for i := pm + 1; i < n; i++ {
		if p[i] < m {
			sum++
		} else {
			sum--
		}
		add(sum)
	}
	// m单独一个的时候
	res := cnt[n+0]
	// 简单点，分开算奇数和偶数长度的结果
	sum = 0
	for i := pm - 1; i >= 0; i-- {
		if p[i] > m {
			sum++
		} else {
			sum--
		}
		res += cnt[sum+n]
	}
	// 偶数长度的时候
	sum = -1
	res += cnt[sum+n]
	for i := pm - 1; i >= 0; i-- {
		if p[i] > m {
			sum++
		} else {
			sum--
		}
		res += cnt[sum+n]
	}

	return res
}

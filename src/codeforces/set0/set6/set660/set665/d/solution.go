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
	a := readNNums(reader, n)
	res := solve(a)

	fmt.Println(len(res))

	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func solve(a []int) []int {
	sort.Ints(a)

	n := len(a)
	ma := a[n-1]
	if ma == 1 {
		return a
	}
	ma *= 2

	var primes []int
	lpf := make([]int, ma+1)
	for i := 2; i <= ma; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j > ma {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}

	var i int
	for i < len(a) && a[i] == 1 {
		i++
	}
	arr := make([]int, i)
	copy(arr, a)
	if i > 0 {
		// 只能再增加一个
		for j := i; j < len(a); j++ {
			if lpf[a[j]+1] == a[j]+1 {
				arr = append(arr, a[j])
				return arr
			}
		}
		if i > 1 {
			return arr
		}
	}
	// len(arr) <= 1
	for j := i; j < len(a); j++ {
		for k := j + 1; k < len(a); k++ {
			if lpf[a[j]+a[k]] == a[j]+a[k] {
				return []int{a[j], a[k]}
			}
		}
	}
	if i > 0 {
		return arr
	}
	return []int{a[0]}
}

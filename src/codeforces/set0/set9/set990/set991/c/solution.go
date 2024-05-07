package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	fmt.Println(solve(n))
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

func solve(n int) int {
	// 显然k超过 n/2， 肯定是可以的
	if n <= 2 {
		return 1
	}
	// 然后直觉上k越大越好
	// 可以这样证明, 假设 k1, k2, k1 < k2
	// 然后考虑bob取的序列 n/k, (n - k) /k2, ...
	// 当k越大时，这个序列越短，且序列中的每个数字，都更小
	half := (n + 1) / 2
	check := func(k int) bool {
		var alice int
		for num := n; num > 0; num -= num / 10 {
			if num > k {
				alice += k
				num -= k
			} else {
				alice += num
				num = 0
			}
		}
		return alice >= half
	}

	l, r := 1, half

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}

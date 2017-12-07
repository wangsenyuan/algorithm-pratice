package main

import (
	"bufio"
	"os"
	"fmt"
)

func readInt(bs []byte, from int, val *int) int {
	tmp := 0
	i := from
	for i < len(bs) && bs[i] != ' ' {
		tmp = tmp*10 + int(bs[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for t > 0 {
		scanner.Scan()
		s := scanner.Bytes()
		var n int
		
		scanner.Scan()
		readInt(scanner.Bytes(), 0, &n)

		filters := make([][]byte, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			filters[i] = scanner.Bytes()
		}

		res := solve(s, n, filters)
		fmt.Println(res)

		t--
	}
}

const MOD = 1000000007

func solve(s []byte, n int, filters [][]byte) int {
	cnts := make([]int, 1024)
	
	for i := 0; i < n; i++ {
		num := toNum(filters[i])
		cnts[num]++
	}

	v := make([]int64, 1024)
	v[0] = 1
	for i := 0; i < 1024; i++ {
		if cnts[i] > 0 {
			for j := 0; j < 1024; j++ {
				if j <= (j ^ i) {
					tmp := ((v[j] + v[j^i]) * pow(2, cnts[i] - 1)) % MOD
					v[j], v[j^i] = tmp, tmp
				}
			}
		}
	}

	return int(v[toNum(s)])
}

func toNum(bs []byte) int {
	num := 0

	for i := 0; i < len(bs); i++ {
		x := 0
		if bs[i] == 'w' || bs[i] == '+' {
			x = 1
		}
		num = num * 2 + x
	}
	//fmt.Println(num)
	return num
}

func pow(a int, b int) int64 {
	if b == 0 {
		return 1
	}
	x := pow(a, b / 2)
	y := (x * x) % MOD
	if b % 2 == 1 {
		return (y * int64(a)) % MOD
	}
	return y
}
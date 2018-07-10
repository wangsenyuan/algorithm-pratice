package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("./A-small-practice.in")
	f, err := os.Open("./B-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	tc := readNum(scanner)

	for i := 1; i <= tc; i++ {
		n := readNum(scanner)
		m := readNum(scanner)
		flavors := make([][]int, m)
		for j := 0; j < m; j++ {
			tmp := readNumsUntilEnd(scanner, 10)
			flavors[j] = tmp[1:]
		}
		res := solve(n, m, flavors)
		if len(res) == 0 {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", i)
			continue
		}
		fmt.Printf("Case #%d: ", i)
		for j, ans := range res {
			if j < len(res)-1 {
				fmt.Printf("%d ", ans)
			} else {
				fmt.Printf("%d\n", ans)
			}
		}
	}

}

func solve(n int, m int, flavors [][]int) []int {
	likes := make([]map[int]int, m)
	malteFlavor := make([]int, m)
	for i, flavor := range flavors {
		likes[i] = make(map[int]int)
		for j := 0; j < len(flavor); j += 2 {
			likes[i][flavor[j]] = flavor[j+1]
			if flavor[j+1] == 1 {
				malteFlavor[i] = flavor[j]
			}
		}
	}

	res := make([]int, n+1)

	needToCheck := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		needToCheck[i] = make([]int, 0, 10)
	}

	stack := make([]int, 0, m)
	var p int
	for i := 0; i < m; i++ {
		stack = append(stack, i)
	}

	for p < len(stack) {
		u := stack[p]
		p++
		var found bool
		for j, x := range likes[u] {
			if x == res[j] {
				if x == 0 {
					needToCheck[j] = append(needToCheck[j], u)
				}
				found = true
				break
			}
		}
		if found {
			continue
		}
		if malteFlavor[u] == 0 || res[malteFlavor[u]] == 1 {
			return nil
		}
		res[malteFlavor[u]] = 1
		stack = append(stack, needToCheck[malteFlavor[u]]...)
		needToCheck[malteFlavor[u]] = needToCheck[malteFlavor[u]][0:0]
	}

	return res[1:]
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func readNumsUntilEnd(scanner *bufio.Scanner, n int) []int {
	res := make([]int, 0, n)
	x := 0
	scanner.Scan()
	bs := scanner.Bytes()
	for x < len(bs) {
		for x < len(bs) && bs[x] == ' ' {
			x++
		}
		var tmp int
		x = readInt(bs, x, &tmp)
		x++
		res = append(res, tmp)
	}
	return res
}

func readFloat64(bs []byte, from int, val *float64) int {
	i := from
	sign := 1
	if bs[i] == '-' {
		sign = -1
		i++
	}

	var dec float64
	for i < len(bs) && bs[i] != '.' && bs[i] != ' ' {
		dec = dec*10 + float64(bs[i]-'0')
		i++
	}
	*val = dec

	if i == len(bs) || bs[i] == ' ' {
		//no fraction
		return i
	}
	i++
	var frac float64
	base := 1.0
	for i < len(bs) && bs[i] != ' ' {
		frac = frac*10 + float64(bs[i]-'0')
		base *= 10
		i++
	}
	*val += frac / base
	return i * sign
}

func readNFloat64s(scanner *bufio.Scanner, n int) []float64 {
	res := make([]float64, n)

	pos := 0
	scanner.Scan()
	bs := scanner.Bytes()
	//fmt.Printf("[debug] %s\n", string(bs))
	for i := 0; i < n; i++ {
		for pos < len(bs) && bs[pos] == ' ' {
			pos++
		}

		pos = readFloat64(bs, pos, &res[i])
	}
	return res
}

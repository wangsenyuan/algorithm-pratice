package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64s(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	pos := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for pos < len(scanner.Bytes()) && scanner.Bytes()[pos] == ' ' {
			pos++
		}
		pos = readInt64(scanner.Bytes(), pos, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	line := readNNums(scanner, 3)
	n, r, c := line[0], line[1], line[2]
	solver := NewSolver(r, c)

	for i := 0; i < n; i++ {
		scanner.Scan()
		step := scanner.Text()
		if step[0] == 'W' {
			solver.ConnectWire(step[1:])
		} else if step[0] == 'V' {
			solver.AddVoltage(step[1:])
		} else if step[0] == 'R' {
			solver.RemoveVoltage(step[1:])
		} else {
			ans := solver.CheckLed(step[1:])
			if ans {
				fmt.Println("ON")
			} else {
				fmt.Println("OFF")
			}
		}
	}
}

type Solver struct {
	r       int
	c       int
	n       int
	arr     []int
	cnt     []int
	sources []int
}

func NewSolver(r int, c int) Solver {
	n := r * c
	arr := make([]int, n)
	cnt := make([]int, n)
	sources := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}

	return Solver{r, c, n, arr, cnt, sources}
}

func (solver *Solver) Find(x int) int {
	if solver.arr[x] != x {
		solver.arr[x] = solver.Find(solver.arr[x])
	}
	return solver.arr[x]
}

func (solver *Solver) ConnectWire(s string) {
	n := len(s)
	px := solver.FindRoot(s[:n/2])
	py := solver.FindRoot(s[n/2:])

	if px == py {
		return
	}
	if solver.cnt[px] >= solver.cnt[py] {
		solver.arr[py] = px
		solver.cnt[px] += solver.cnt[py]
		solver.sources[px] += solver.sources[py]
	} else {
		solver.arr[px] = py
		solver.cnt[py] += solver.cnt[px]
		solver.sources[py] += solver.sources[px]
	}
}

func (solver *Solver) FindRoot(s string) int {
	coord := parseAsNums(s, 2)
	a, b := coord[0], coord[1]
	a--
	b--

	b /= 5

	x := b*solver.c + a
	return solver.Find(x)
}

func (solver *Solver) AddVoltage(s string) {
	px := solver.FindRoot(s)
	solver.sources[px]++
}

func (solver *Solver) RemoveVoltage(s string) {
	px := solver.FindRoot(s)
	solver.sources[px]--
}

func (solver *Solver) CheckLed(s string) bool {
	n := len(s)
	px := solver.FindRoot(s[:n/2])
	py := solver.FindRoot(s[n/2:])

	return (solver.sources[px] > 0) != (solver.sources[py] > 0)
}

func parseAsNums(s string, cnt int) []int {
	res := make([]int, cnt)

	n := len(s)
	x := n / cnt

	for i := 0; i < cnt; i++ {
		var num int

		for j := 0; j < x; j++ {
			k := i*cnt + j
			if s[k] >= 'A' && s[k] <= 'Z' {
				num = num*52 + int(s[k]-'A')
			} else {
				num = num*52 + int(s[k]-'a') + 26
			}
		}

		res[i] = num
	}

	return res
}

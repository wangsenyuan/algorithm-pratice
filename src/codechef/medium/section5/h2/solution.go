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
	persons := make([][]int, n)

	for i := 0; i < n; i++ {
		persons[i] = readNNums(reader, 2)
	}

	res := solve(n, persons)

	fmt.Printf("%.7f\n", res)
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(n int, people [][]int) float64 {
	// strength >= 50 & <= 100
	// weight >= 50 & <= 100
	// so total unique pair would be <= 50 * 50
	pairs := make(map[Pair]int)
	for _, cur := range people {
		pairs[Pair{cur[1], cur[0]}]++
	}
	arr := make([]Pair, 0, len(pairs))
	for cur := range pairs {
		arr = append(arr, cur)
	}

	arr2 := make([]Pair, n)

	check := func(speed float64) bool {
		comp := Comp{speed, arr}

		sort.Sort(&comp)

		var p int

		for i := 0; i < len(arr); i++ {
			for j := 0; j < pairs[arr[i]]; j++ {
				arr2[p] = arr[i]
				p++
			}
		}

		for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
			a, b := arr2[i], arr2[j]
			if float64(a.first+b.first)-speed*float64(a.second+b.second+20) < 0 {
				return false
			}
		}
		return true
	}

	var lo, hi float64 = 0, 2.0

	for i := 0; i < 30; i++ {
		mid := (lo + hi) / 2
		if check(mid) {
			lo = mid
		} else {
			hi = mid
		}
	}

	return (lo + hi) / 2
}

type Pair struct {
	first  int
	second int
}

func (p Pair) Value(speed float64) float64 {
	return float64(p.first) - speed*float64(p.second)
}

type Comp struct {
	speed float64
	arr   []Pair
}

func (this *Comp) Len() int {
	return len(this.arr)
}

func (this *Comp) Less(i, j int) bool {
	a, b := this.arr[i], this.arr[j]
	return a.Value(this.speed) < b.Value(this.speed)
}

func (this *Comp) Swap(i, j int) {
	this.arr[i], this.arr[j] = this.arr[j], this.arr[i]
}

package main

import (
	"bufio"
	"fmt"
	"sort"
)

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

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	res := readNNums(scanner, 3)
	a, b, c = res[0], res[1], res[2]
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	var n int64
	var m, k int

	fmt.Scanf("%d %d %d", &n, &m, &k)
	edges := make([][]int64, k)

	for i := 0; i < k; i++ {
		edges[i] = make([]int64, 4)
		fmt.Scanf("%d %d %d %d", &edges[i][0], &edges[i][1], &edges[i][2], &edges[i][3])
	}

	fmt.Println(solve(n, m, edges))
}

const MOD = 1000000007

func pow(a int, b int64) int64 {
	A := int64(a)
	R := int64(1)

	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}

	return R
}

func solve(n int64, m int, edges [][]int64) int {
	froms := make(map[Pair][]Pair)

	events := make([]Event, 0, 2*len(edges)+1)

	for _, edge := range edges {
		x1, y1 := edge[0], int(edge[1])
		x2, y2 := edge[2], int(edge[3])

		events = append(events, NewEvent(x1, y1, false))
		events = append(events, NewEvent(x2, y2, true))

		to := Pair{x2, y2}
		from := Pair{x1, y1}
		if _, found := froms[to]; found {
			froms[to] = append(froms[to], from)
		} else {
			froms[to] = []Pair{from}
		}
	}

	events = append(events, NewEvent(n+1, 0, false))

	sort.Sort(Events(events))

	events = unique(events)

	M := int64(m)

	var curPos int64
	var ans int64 = 1
	var sum int64 = 1

	mem := make(map[Pair]int64)

	for i := 0; i < len(events); i++ {
		if curPos < events[i].x {
			ans = (sum * pow(m, events[i].x-curPos-1)) % MOD
			curPos = events[i].x
			sum = (ans * M) % MOD
		}

		to := Pair{events[i].x, events[i].y}

		if events[i].end {
			for _, from := range froms[to] {
				sum += mem[from]
				sum %= MOD
				if zz, found := mem[to]; found {
					mem[to] = (zz + mem[from]) % MOD
				}
			}
		} else {
			if _, found := mem[to]; !found {
				mem[to] = ans
			}
		}
	}

	return int(mem[Pair{n + 1, 0}])
}

func unique(events []Event) []Event {
	var j int

	for i := 1; i <= len(events); i++ {
		if i == len(events) || events[i-1].less(events[i]) {
			events[j] = events[i-1]
			j++
		}
	}
	return events[:j]
}

type Pair struct {
	first  int64
	second int
}

type Event struct {
	x   int64
	y   int
	end bool
}

func (this Event) less(that Event) bool {
	if this.x < that.x {
		return true
	}

	if this.x == that.x {
		if this.end != that.end {
			// make sure in first, out last
			return !this.end
		}
		// i.end == j.end
		return this.y < that.y
	}

	return false
}

func NewEvent(x int64, y int, end bool) Event {
	return Event{x, y, end}
}

type Events []Event

func (this Events) Len() int {
	return len(this)
}

func (this Events) Less(i, j int) bool {
	return this[i].less(this[j])
}

func (this Events) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

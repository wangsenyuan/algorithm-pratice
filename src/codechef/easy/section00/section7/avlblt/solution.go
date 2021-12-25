package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
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
	scanner := bufio.NewScanner(os.Stdin)

	T := readNum(scanner)

	blocks := make([][]string, 7)
	for i := 0; i < 7; i++ {
		n := readNum(scanner)

		blocks[i] = make([]string, n)

		for j := 0; j < n; j++ {
			scanner.Scan()
			blocks[i][j] = scanner.Text()
		}
	}

	res := solve(T, blocks)

	for i := 0; i < 7; i++ {
		fmt.Println(len(res[i]))
		for j := 0; j < len(res[i]); j++ {
			fmt.Println(res[i][j])
		}
	}
}

const DAY_MINS = 24 * 60

func solve(T int, blocks [][]string) [][]string {
	clocks := make([][]*clock, 7)
	for i := 0; i < 7; i++ {
		block := blocks[i]
		clocks[i] = make([]*clock, 0, len(block))

		for _, s := range block {
			clocks[i] = append(clocks[i], parseString(s))
		}
	}

	M := T * 60

	res := make([][]*clock, 7)

	for i := 0; i < 7; i++ {
		res[i] = make([]*clock, 0, 3)
	}

	for i := 0; i < 7; i++ {
		clks := clocks[i]
		for _, clk := range clks {
			x := clk.start - M
			y := clk.end - M
			if y <= x {
				y += DAY_MINS
			}
			if x >= 0 {
				// y should also > 0, in the same day
				res[i] = append(res[i], newClock(x, y))
			} else if y >= 0 {
				//split the clock
				res[i] = append(res[i], newClock(0, y))
				j := (i + 6) % 7
				res[j] = append(res[j], newClock(x+DAY_MINS, DAY_MINS))
			} else {
				j := (i + 6) % 7
				res[j] = append(res[j], newClock(x+DAY_MINS, y+DAY_MINS))
			}
		}
	}

	ans := make([][]string, 7)
	for i := 0; i < 7; i++ {
		ans[i] = make([]string, 0, len(res[i]))
		sort.Sort(ClocksSlice(res[i]))
		var j = 0
		for k := 1; k <= len(res[i]); k++ {
			if k == len(res[i]) || res[i][k].start > res[i][k-1].end {
				x := res[i][j].start
				y := res[i][k-1].end
				ans[i] = append(ans[i], formatTime(x, y))
				j = k
			}
		}
	}

	return ans
}

type clock struct {
	start int
	end   int
}

func newClock(start, end int) *clock {
	clk := new(clock)
	clk.start = start
	clk.end = end
	return clk
}

func parseString(s string) *clock {
	clk := new(clock)

	ss := strings.Split(s, " ")
	clk.start = parseTime(ss[0])
	clk.end = parseTime(ss[1])

	return clk
}

func parseTime(s string) int {
	ss := strings.Split(s, ":")
	hour, _ := strconv.Atoi(ss[0])
	min, _ := strconv.Atoi(ss[1])

	return hour*60 + min
}

func formatTime(start int, end int) string {
	x := formatTime2(start)
	y := formatTime2(end)
	return x + " " + y
}

func formatTime2(time int) string {
	time = time % DAY_MINS
	hour := time / 60
	min := time % 60
	return fmt.Sprintf("%02d:%02d", hour, min)
}

type ClocksSlice []*clock

func (this ClocksSlice) Len() int {
	return len(this)
}

func (this ClocksSlice) Less(i, j int) bool {
	if this[i].start < this[j].start {
		return true
	}

	if this[i].start == this[j].start && this[i].end < this[j].end {
		return true
	}

	return false
}

func (this ClocksSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

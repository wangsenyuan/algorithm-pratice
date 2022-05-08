package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	S := make([]string, n)

	for i := 0; i < n; i++ {
		S[i] = readString(reader)
	}

	res := solve(n, S)

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, commands []string) []int {
	ans := make([]int, n)

	var x, y int64

	strips := make([]Stripe, n+1)
	strips[0] = Stripe{0, 0, 0, 0}

	for i := 0; i < n; i++ {
		cmd := parseCommand(commands[i])

		if cmd.dir == 'L' {
			strips[i+1] = NewStripe(x-1, y, x-cmd.dis, y)
			x -= cmd.dis
		} else if cmd.dir == 'R' {
			strips[i+1] = NewStripe(x+1, y, x+cmd.dis, y)
			x += cmd.dis
		} else if cmd.dir == 'U' {
			strips[i+1] = NewStripe(x, y+1, x, y+cmd.dis)
			y += cmd.dis
		} else {
			strips[i+1] = NewStripe(x, y-1, x, y-cmd.dis)
			y -= cmd.dis
		}

		var horizontal bool
		if cmd.dir == 'L' || cmd.dir == 'R' {
			horizontal = true
		}
		var ps []Point
		for j := 0; j <= i; j++ {
			ok, intersection := intersect(strips[i+1], strips[j])

			if ok {
				if horizontal {
					ps = append(ps, Point{intersection.x1 - strips[i].x1, false})
					ps = append(ps, Point{intersection.x2 - strips[i].x1 + 1, true})
				} else {
					ps = append(ps, Point{intersection.y1 - strips[i].y1, false})
					ps = append(ps, Point{intersection.y2 - strips[i].y1 + 1, true})
				}
			}
		}

		sort.Sort(Points(ps))

		var covered int
		var last int64
		var painted = cmd.dis

		for j := 0; j < len(ps); j++ {
			p := ps[j]
			if covered > 0 {
				painted -= p.coord - last
			}
			if p.begin {
				covered--
			} else {
				covered++
			}
			last = p.coord
		}
		ans[i] = int(painted)
	}

	return ans
}

type Stripe struct {
	x1, y1, x2, y2 int64
}

func NewStripe(x1, y1, x2, y2 int64) Stripe {
	return Stripe{min(x1, x2), min(y1, y2), max(x1, x2), max(y1, y2)}
}

func intersect(a, b Stripe) (bool, Stripe) {
	x1 := max(a.x1, b.x1)
	y1 := max(a.y1, b.y1)
	x2 := min(a.x2, b.x2)
	y2 := min(a.y2, b.y2)
	if x1 > x2 || y1 > y2 {
		return false, Stripe{}
	}
	return true, Stripe{x1, y1, x2, y2}
}

type Point struct {
	coord int64
	begin bool
}

func (this Point) Less(that Point) bool {
	if this.coord != that.coord {
		return this.coord < that.coord
	}
	if this.begin != that.begin {
		return !this.begin
	}
	return false
}

type Points []Point

func (this Points) Len() int {
	return len(this)
}

func (this Points) Less(i, j int) bool {
	return this[i].Less(this[j])
}

func (this Points) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

type Command struct {
	dir byte
	dis int64
}

func parseCommand(s string) *Command {
	dir := s[0]
	var dis int
	readInt([]byte(s), 2, &dis)
	return &Command{dir, int64(dis)}
}

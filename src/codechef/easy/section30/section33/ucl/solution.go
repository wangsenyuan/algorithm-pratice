package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		records := make([]string, 12)
		for i := 0; i < 12; i++ {
			records[i] = readString(reader)
		}
		res := solve(records)
		buf.WriteString(fmt.Sprintf("%s %s\n", res[0], res[1]))
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
		if s[i] == '\n' || s[i] == '\r' {
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

func solve(records []string) []string {
	scores := make(map[string]int)
	diffs := make(map[string]int)

	for _, record := range records {
		tms := ParseRecord(record)
		a, b := tms[0], tms[1]
		if a.goals > b.goals {
			scores[a.name] += 3
		} else if a.goals == b.goals {
			scores[a.name] += 1
			scores[b.name] += 1
		} else {
			scores[b.name] += 3
		}
		diffs[a.name] += a.goals - b.goals
		diffs[b.name] += b.goals - a.goals
	}

	res := make([]Result, 0, len(scores))

	for k, v := range scores {
		d := diffs[k]
		res = append(res, Result{k, v, d})
	}

	sort.Sort(Results(res))

	return []string{res[0].name, res[1].name}
}

type Result struct {
	name   string
	scores int
	diffs  int
}

func (this Result) Less(that Result) bool {
	return this.scores > that.scores || this.scores == that.scores && this.diffs > that.diffs
}

type Results []Result

func (this Results) Len() int {
	return len(this)
}

func (this Results) Less(i, j int) bool {
	return this[i].Less(this[j])
}

func (this Results) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type Team struct {
	name  string
	goals int
}

func ParseRecord(record string) []Team {
	ss := strings.Split(record, " ")
	host := ss[0]
	var hostGoals int
	readInt([]byte(ss[1]), 0, &hostGoals)
	away := ss[4]
	var awayGoals int
	readInt([]byte(ss[3]), 0, &awayGoals)

	return []Team{{host, hostGoals}, {away, awayGoals}}
}

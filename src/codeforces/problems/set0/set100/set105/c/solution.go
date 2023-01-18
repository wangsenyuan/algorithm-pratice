package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	items := make([]string, n)
	for i := 0; i < n; i++ {
		items[i] = readString(reader)
	}
	m := readNum(reader)
	residents := make([]string, m)
	for i := 0; i < m; i++ {
		residents[i] = readString(reader)
	}
	res := solve(items, residents)

	for _, row := range res {
		fmt.Println(row)
	}
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func solve(items []string, residents []string) []string {
	its := mapAs(items, NewItem)
	rds := mapAs(residents, NewResident)

	cnt := make(map[string]int)

	for _, r := range rds {
		cnt[r.home]++
	}

	free := false
	for _, it := range its {
		if cnt[it.name] < it.size {
			free = true
			break
		}
	}
	if !free {
		return solve1(its, rds)
	}
	return solve2(its, rds)
}

func solve2(items []Item, residents []Resident) []string {

	sort.Slice(residents, func(i, j int) bool {
		return residents[i].bonus > residents[j].bonus
	})
	// can move as wanted

	getByType := func(tpe string, sz int) (int, []string) {
		var names []string
		var sum int
		var cnt int
		for _, r := range residents {
			if r.tpe == tpe && cnt < sz {
				cnt++
				sum += r.bonus
				names = append(names, r.name)
			}
		}
		return sum, names
	}

	tps := map[string]string{
		"weapon": "gladiator",
		"armor":  "sentry",
		"orb":    "physician",
	}

	score := make(map[string]Pair)

	for i, it := range items {
		rs, _ := getByType(tps[it.clazz], it.size)
		if it.clazz == "weapon" {
			rs += it.atk
		} else if it.clazz == "armor" {
			rs += it.def
		} else {
			rs += it.res
		}
		if v, ok := score[it.clazz]; !ok || v.first < rs {
			score[it.clazz] = Pair{rs, i}
		}
	}
	// var res []string
	type Ans struct {
		name string
		rs   []string
		sz   int
	}

	xx := make(map[string]*Ans)

	used := make(map[string]bool)

	for k, v := range score {
		i := v.second
		_, names := getByType(tps[items[i].clazz], items[i].size)

		for _, name := range names {
			used[name] = true
		}
		// res = append(res, CreateAnswer(items[i].name, names))
		xx[k] = &Ans{items[i].name, names, items[i].size}
	}

	for _, rd := range residents {
		if used[rd.name] {
			continue
		}
		for k := range xx {
			if len(xx[k].rs) < xx[k].sz {
				xx[k].rs = append(xx[k].rs, rd.name)
				break
			}
		}
	}

	var res []string
	wps := xx["weapon"].rs
	res = append(res, CreateAnswer(xx["weapon"].name, wps))

	defs := xx["armor"].rs
	res = append(res, CreateAnswer(xx["armor"].name, defs))

	ress := xx["orb"].rs
	res = append(res, CreateAnswer(xx["orb"].name, ress))

	return res
}

func solve1(items []Item, residents []Resident) []string {
	tps := map[string]string{
		"weapon": "gladiator",
		"armor":  "sentry",
		"orb":    "physician",
	}

	// can't move
	getByHome := func(home string, tp string) (int, []string) {
		var names []string
		var sum int
		for _, r := range residents {
			if r.home == home {
				if r.tpe == tp {
					sum += r.bonus
				}
				names = append(names, r.name)
			}
		}
		return sum, names
	}

	score := make(map[string]Pair)

	for i, it := range items {
		rs, _ := getByHome(it.name, tps[it.clazz])
		if it.clazz == "weapon" {
			rs += it.atk
		} else if it.clazz == "armor" {
			rs += it.def
		} else {
			rs += it.res
		}
		if v, ok := score[it.clazz]; !ok || v.first < rs {
			score[it.clazz] = Pair{rs, i}
		}
	}

	var res []string

	_, wps := getByHome(items[score["weapon"].second].name, "gladiator")
	res = append(res, CreateAnswer(items[score["weapon"].second].name, wps))

	_, defs := getByHome(items[score["armor"].second].name, "sentry")
	res = append(res, CreateAnswer(items[score["armor"].second].name, defs))

	_, ress := getByHome(items[score["orb"].second].name, "physician")
	res = append(res, CreateAnswer(items[score["orb"].second].name, ress))

	return res
}

func CreateAnswer(name string, ss []string) string {
	var buf bytes.Buffer
	buf.WriteString(name)
	buf.WriteString(fmt.Sprintf(" %d", len(ss)))
	sort.Strings(ss)
	for _, s := range ss {
		buf.WriteByte(' ')
		buf.WriteString(s)
	}
	return buf.String()
}

type Pair struct {
	first  int
	second int
}

type Item struct {
	name  string
	clazz string
	atk   int
	def   int
	res   int
	size  int
}

func NewItem(s string) Item {
	ss := strings.Split(s, " ")

	name := ss[0]
	clazz := ss[1]
	atk, _ := strconv.Atoi(ss[2])
	def, _ := strconv.Atoi(ss[3])
	res, _ := strconv.Atoi(ss[4])
	size, _ := strconv.Atoi(ss[5])

	return Item{name, clazz, atk, def, res, size}
}

type Resident struct {
	name  string
	tpe   string
	bonus int
	home  string
}

func NewResident(s string) Resident {
	ss := strings.Split(s, " ")
	name := ss[0]
	tpe := ss[1]
	bonus, _ := strconv.Atoi(ss[2])
	home := ss[3]
	return Resident{name, tpe, bonus, home}
}

type ItemOrResident interface {
	Item | Resident
}

func mapAs[T ItemOrResident](arr []string, f func(string) T) []T {
	res := make([]T, len(arr))
	for i, s := range arr {
		res[i] = f(s)
	}
	return res
}

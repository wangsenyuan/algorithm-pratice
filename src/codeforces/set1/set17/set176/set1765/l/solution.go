package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	workers := make([]string, n)
	for i := 0; i < n; i++ {
		workers[i] = readString(reader)[2:]
	}
	holidays := readNNums(reader, m)
	projects := make([][]int, k)
	for i := 0; i < k; i++ {
		s, _ := reader.ReadBytes('\n')
		var pos, cnt int
		pos = readInt(s, 0, &cnt)
		projects[i] = make([]int, cnt)
		for j := 0; j < cnt; j++ {
			pos = readInt(s, pos+1, &projects[i][j])
		}
	}
	res := solve(workers, holidays, projects)
	var buf bytes.Buffer

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

var days map[string]int

func init() {
	days = make(map[string]int)
	days["Monday"] = 0
	days["Tuesday"] = 1
	days["Wednesday"] = 2
	days["Thursday"] = 3
	days["Friday"] = 4
	days["Saturday"] = 5
	days["Sunday"] = 6
}

func solve(workers []string, holidays []int, projects [][]int) []int {
	ws := make([]int, len(workers))
	wk := make([]*IntHeap, len(workers))
	for i := 0; i < len(workers); i++ {
		s := strings.Split(workers[i], " ")

		for _, x := range s {
			ws[i] |= (1 << days[x])
		}

		pq := make(IntHeap, 0, 1)
		wk[i] = &pq
	}

	cur := make([]map[int]int, 7)
	for i := 0; i < 7; i++ {
		cur[i] = make(map[int]int)
	}

	ans := make([]int, len(projects))
	last := make([]int, len(projects))

	for id, a := range projects {
		ans[id] = -1
		for i := 0; i < len(a); i++ {
			a[i]--
		}
		for j := 0; j < 7; j++ {
			if working(ws[a[0]], j) {
				cur[j][a[0]]++
			}
		}
		heap.Push(wk[a[0]], id)
	}

	var done int

	now := make([]int, len(workers))
	sv := make([]int, len(workers))
	for j, d := 0, 1; done < len(projects); d++ {
		if j < len(holidays) && holidays[j] == d {
			j++
			continue
		}
		wd := (d - 1) % 7
		var it int
		for x := range cur[wd] {
			now[it] = x
			it++
		}

		for i := 0; i < it; i++ {
			x := now[i]
			for ii := 0; ii < 7; ii++ {
				if cur[ii][x] == 1 {
					delete(cur[ii], x)
				} else if cur[ii][x] > 1 {
					cur[ii][x]--
				}
			}
			sv[i] = heap.Pop(wk[x]).(int)
		}

		for i := 0; i < it; i++ {
			y := sv[i]
			last[y]++
			if last[y] == len(projects[y]) {
				ans[y] = d
				done++
				continue
			}
			wid := projects[y][last[y]]
			heap.Push(wk[wid], y)
			for ii := 0; ii < 7; ii++ {
				if working(ws[wid], ii) {
					cur[ii][wid]++
				}
			}
		}
	}

	return ans
}

func working(flag int, i int) bool {
	return (flag>>i)&1 == 1
}

func solve1(workers []string, holidays []int, projects [][]int) []int {
	// 是不是按天去模拟就可以了？
	// 每个project，都排在它当前任务需要的工人的队伍里
	// 每个人都排在下一个能够work的天数里
	workDays := make(map[int][]int)

	pq := make(IntHeap, 0, 1)

	addActive := func(workerId int, day int) {
		if workDays[day] == nil {
			workDays[day] = make([]int, 0, 1)
			heap.Push(&pq, day)
		}
		workDays[day] = append(workDays[day], workerId)
	}

	ws := make([]*Worker, len(workers))
	tasks := make(map[int]*IntHeap)
	for i := 0; i < len(workers); i++ {
		ws[i] = NewWorker(i, workers[i])
		addActive(i, ws[i].workDays[0])
		t := make(IntHeap, 0, 1)
		tasks[i] = &t
	}

	pos := make([]int, len(projects))
	for i, cur := range projects {
		workerId := cur[0] - 1
		heap.Push(tasks[workerId], i)
	}

	ans := make([]int, len(projects))

	sort.Ints(holidays)
	for i := 0; i < len(holidays); i++ {
		holidays[i]--
	}

	cnt := len(projects)

	pending := make([]int, len(projects))

	var hi int

	for pq.Len() > 0 && cnt > 0 {
		now := heap.Pop(&pq).(int)

		var needBreak bool

		for hi < len(holidays) && holidays[hi] <= now {
			needBreak = holidays[hi] == now
			hi++
		}

		ids := workDays[now]

		var it int

		for _, id := range ids {
			worker := ws[id]

			if !needBreak {
				// 最优先的任务
				if tasks[id].Len() > 0 {
					pid := heap.Pop(tasks[id]).(int)
					project := projects[pid]
					pos[pid]++
					if pos[pid] == len(project) {
						// this project done
						ans[pid] = now + 1
						cnt--
					} else {
						// pending to next worker
						pending[it] = pid
						it++
					}
				}
			}

			worker.workDays[worker.index] += 7
			worker.index = (worker.index + 1) % len(worker.workDays)
			addActive(worker.id, worker.workDays[worker.index])
		}

		delete(workDays, now)

		for i := 0; i < it; i++ {
			pid := pending[i]
			cur := projects[pid]
			wid := cur[pos[pid]] - 1
			heap.Push(tasks[wid], pid)
		}
	}

	return ans
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Worker struct {
	id       int
	workDays []int
	index    int
}

func NewWorker(id int, s string) *Worker {
	ss := strings.Split(s, " ")
	var d []int
	for _, x := range ss {
		d = append(d, days[x])
	}
	return &Worker{id, d, 0}
}

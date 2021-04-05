package e

import "sort"

const INF = 1000000007

func processTasks(tasks [][]int) int {
	events := make([]Event, 2*len(tasks))

	for i := 0; i < len(tasks); i++ {
		events[2*i] = Event{i, 0, tasks[i][0]}
		events[2*i+1] = Event{i, 1, tasks[i][1] + 1}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Less(events[j])
	})

	pp := make([]int, len(events))
	var n int
	ii := make(map[int]int)
	for i := 1; i <= len(events); i++ {
		if i == len(events) || events[i].pos > events[i-1].pos {
			pp[n] = events[i-1].pos
			n++
			ii[events[i-1].pos] = n
		}
	}

	all := make([]int, n)
	use := make([]int, n)
	for i := 1; i < n; i++ {
		all[i] = pp[i] - pp[i-1]
	}

	bit := make([]int, n+1)
	add := func(p int, v int) {
		//p++
		for p <= n {
			bit[p] += v
			p += p & -p
		}
	}

	count := func(p int) int {
		//p++
		var res int

		for p > 0 {
			res += bit[p]
			p -= p & -p
		}
		return res
	}

	set := make([]int, n)
	for i := 0; i < n; i++ {
		set[i] = i
	}

	var find func(x int) int

	find = func(x int) int {
		if set[x] != x {
			set[x] = find(set[x])
		}
		return set[x]
	}

	for i := 0; i < len(events); i++ {
		cur := events[i]
		if cur.tpe == 0 {
			continue
		}
		task := tasks[cur.id]
		l, r, need := task[0], task[1]+1, task[2]
		l = ii[l]
		r = ii[r]
		has := count(r-1) - count(l-1)
		if has >= need {
			continue
		}
		pos := set[r-1]
		need -= has
		for need > 0 {
			left := all[pos] - use[pos]
			if left > need {
				use[pos] += need
				add(pos, need)
				break
			}
			if left == need {
				use[pos] += need
				add(pos, need)
				set[pos] = set[pos-1]
				break
			}
			// left < need
			use[pos] += left
			need -= left
			add(pos, left)
			set[pos] = set[pos-1]
			pos = find(pos)
		}
	}
	return count(n - 1)
}

type Event struct {
	id  int
	tpe int
	pos int
}

func (this Event) Less(that Event) bool {
	return this.pos < that.pos
}

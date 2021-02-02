package main

import (
	"fmt"
	"sort"
	"strings"
)

var LABELS = [26]string{}

func init() {
	for i := 0; i < 26; i++ {
		LABELS[i] = fmt.Sprint('A' + i)
	}
}

func main() {
	t := 0
	fmt.Scanf("%d", &t)
	for i := 1; i <= t; i++ {
		n := 0
		fmt.Scanf("%d", &n)
		parties := make([]*Party, 0, n)
		for j := 0; j < n; j++ {
			x := 0
			fmt.Scanf("%d", &x)
			p := new(Party)
			p.cnt = x
			p.label = LABELS[j]
			parties = append(parties, p)
		}
		sort.Sort(sort.Reverse(byCnt(parties)))
		//fmt.Printf("after sort %s\n", parties)
		plan := play(parties, make([]string, 0, n))
		fmt.Printf("Case #%d: %s\n", i, strings.Join(plan, " "))
	}
}

type Party struct {
	cnt   int
	label string
}

func (p *Party) String() string {
	return fmt.Sprintf("{%d %s}", p.cnt, p.label)
}

func (p *Party) copy(cnt int) *Party {
	that := new(Party)
	that.cnt = cnt
	that.label = p.label
	return that
}

func sortParty(a *Party, parties []*Party) []*Party {
	if a.cnt == 0 {
		return parties
	}

	if len(parties) == 0 {
		return []*Party{a}
	}

	h := parties[0]
	if a.cnt >= h.cnt {
		return append([]*Party{a}, parties...)
	}

	return append([]*Party{h}, sortParty(a, parties[1:])...)
}

func play(parties []*Party, plan []string) []string {
	//fmt.Printf("%s ---- %s\n", parties, plan)
	if len(parties) == 0 {
		return plan
	}

	if len(parties) == 2 {
		return play2(parties[0], parties[1], plan)
	}

	a, b := parties[0], parties[1]

	if a.cnt-b.cnt >= 2 {
		return play(sortParty(a.copy(a.cnt-2), parties[1:]), updatePlan(plan, a.label+a.label))
	}

	return play(sortParty(a.copy(a.cnt-1), parties[1:]), updatePlan(plan, a.label))
}

func play2(a, b *Party, plan []string) []string {
	//fmt.Printf("%s, %s -- %s\n", a, b, plan)
	if a.cnt == b.cnt && a.cnt == 0 {
		return plan
	}

	if a.cnt == b.cnt {
		return play2(a.copy(a.cnt-1), b.copy(b.cnt-1), updatePlan(plan, a.label+b.label))
	}

	if a.cnt-b.cnt >= 2 {
		return play2(a.copy(a.cnt-2), b, updatePlan(plan, a.label))
	}

	return play2(a.copy(a.cnt-1), b, updatePlan(plan, a.label))
}

func updatePlan(plan []string, x string) []string {
	if len(plan)+1 >= cap(plan) {
		newPlan := make([]string, len(plan), 2*len(plan))
		copy(newPlan, plan)
		//fmt.Printf("%s copied from %s\n", newPlan, plan)
		// return append(newPlan, x)
		plan = newPlan
	}
	return append(plan, x)
}

type byCnt []*Party

func (ps byCnt) Len() int {
	return len(ps)
}

func (ps byCnt) Less(i, j int) bool {
	return ps[i].cnt <= ps[j].cnt
}

func (ps byCnt) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

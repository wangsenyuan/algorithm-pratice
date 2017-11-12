package main

import (
	"sort"
	"fmt"
)

func main() {
	nums := []int{5, 4, 3, 2, 1}
	fmt.Println(findRelativeRanks(nums))
}

func findRelativeRanks(nums []int) []string {
	n := len(nums)

	players := make([]*Player, n)
	for i := 0; i < n; i++ {
		players[i] = &Player{nums[i], i}
	}

	sort.Sort(Players(players))

	res := make([]string, n)

	for i := 0; i < n; i++ {
		res[players[i].idx] = rank(i)
	}

	return res
}
func rank(i int) string {
	if i == 0 {
		return "Gold Medal"
	}

	if i == 1 {
		return "Silver Medal"
	}

	if i == 2 {
		return "Bronze Medal"
	}
	return fmt.Sprintf("%d", i+1)
}

type Player struct {
	score int
	idx   int
}

type Players []*Player

func (p Players) Len() int {
	return len(p)
}

func (p Players) Less(i, j int) bool {
	return p[i].score > p[j].score
}

func (p Players) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

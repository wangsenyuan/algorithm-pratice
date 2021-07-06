package p1921

import "sort"

func eliminateMaximum(dist []int, speed []int) int {
	// 保证每分钟都能杀死一个怪物
	mm := make([]Monster, len(dist))

	for i := 0; i < len(dist); i++ {
		mm[i] = Monster{dist[i], speed[i]}
	}

	sort.Slice(mm, func(i, j int) bool {
		return mm[i].Less(mm[j])
	})

	var cur int
	for i := 0; i < len(mm); i++ {
		if cur*mm[i].speed >= mm[i].dist {
			break
		}
		cur++
	}
	return cur
}

type Monster struct {
	dist  int
	speed int
}

func (this Monster) Less(that Monster) bool {
	// more earlier arrive at city
	// this.dist / this.speed < that.dist / that.speed
	return this.dist*that.speed < that.dist*this.speed
}

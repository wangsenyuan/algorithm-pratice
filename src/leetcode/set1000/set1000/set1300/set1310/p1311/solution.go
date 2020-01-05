package p1311

import "sort"

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	// first find level k friends

	n := len(watchedVideos)
	dist := make([]int, n)

	for i := 0; i < n; i++ {
		dist[i] = -1
	}
	dist[id] = 0

	que := make([]int, n)
	var front, end int
	que[end] = id
	end++

	for front < end {
		u := que[front]
		front++

		for _, v := range friends[u] {
			if dist[v] < 0 {
				dist[v] = dist[u] + 1
				que[end] = v
				end++
			}
		}
	}

	cnt := make(map[string]int)

	for i := 0; i < n; i++ {
		if dist[i] == level {
			for _, video := range watchedVideos[i] {
				cnt[video]++
			}
		}
	}

	ps := make([]Pair, 0, len(cnt))
	for k, v := range cnt {
		ps = append(ps, Pair{k, v})
	}

	sort.Sort(Pairs(ps))

	res := make([]string, len(ps))

	for i := 0; i < len(ps); i++ {
		res[i] = ps[i].first
	}

	return res
}

type Pair struct {
	first  string
	second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].second < ps[j].second || (ps[i].second == ps[j].second && ps[i].first < ps[j].first)
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

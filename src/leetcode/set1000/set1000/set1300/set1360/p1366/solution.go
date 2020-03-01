package p1366

import (
	"bytes"
	"sort"
)

func rankTeams(votes []string) string {
	n := len(votes[0])

	teams := make([]*Team, 26)

	for i := 0; i < 26; i++ {
		team := new(Team)
		team.label = byte(i + 'A')
		team.rank = make([]int, n)
		teams[i] = team
	}

	vis := make([]bool, 26)

	for _, s := range votes {
		for i := 0; i < n; i++ {
			x := int(s[i] - 'A')
			vis[x] = true
			teams[x].rank[i]++
		}
	}

	sort.Sort(Teams(teams))

	var buf bytes.Buffer

	for i := 0; i < 26; i++ {
		j := int(teams[i].label - 'A')
		if !vis[j] {
			continue
		}
		buf.WriteByte(teams[i].label)
	}

	return buf.String()
}

type Team struct {
	label byte
	rank  []int
}

type Teams []*Team

func (teams Teams) Len() int {
	return len(teams)
}

func (teams Teams) Less(i, j int) bool {
	a := teams[i]
	b := teams[j]

	for k := 0; k < len(a.rank); k++ {
		if a.rank[k] > b.rank[k] {
			return true
		} else if a.rank[k] < b.rank[k] {
			return false
		}
	}

	return a.label < b.label
}

func (teams Teams) Swap(i, j int) {
	teams[i], teams[j] = teams[j], teams[i]
}

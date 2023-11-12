package p2933

import (
	"sort"
	"strconv"
)

func findHighAccessEmployees(access_times [][]string) []string {
	sort.Slice(access_times, func(i, j int) bool {
		return access_times[i][1] < access_times[j][1]
	})
	added := make(map[string]bool)
	var res []string
	cnt := make(map[string]int)

	for l, r := 0, 0; r < len(access_times); r++ {
		for l < r && !withInOneHour(access_times[l][1], access_times[r][1]) {
			cnt[access_times[l][0]]--
			l++
		}
		cnt[access_times[r][0]]++
		if cnt[access_times[r][0]] >= 3 && !added[access_times[r][0]] {
			res = append(res, access_times[r][0])
			added[access_times[r][0]] = true
		}
	}

	return res
}

func withInOneHour(a, b string) bool {
	x := parseAsTime(a)
	y := parseAsTime(b)
	return y-x < 60
}

func parseAsTime(s string) int {
	hour, _ := strconv.Atoi(s[:2])
	min, _ := strconv.Atoi(s[2:])
	return hour*60 + min
}

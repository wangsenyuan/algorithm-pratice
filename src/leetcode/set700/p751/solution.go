package p751

import (
	"fmt"
	"strconv"
	"strings"
)

func ipToCIDR(ip string, rang int) []string {
	var x int64
	parts := strings.Split(ip, ".")

	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		x = x*256 + int64(num)
	}
	ans := make([]string, 0, rang)
	for rang > 0 {
		step := x & -x
		for step > int64(rang) {
			step /= 2
		}
		ans = append(ans, longToIP(x, int(step)))
		x += step
		rang -= int(step)
	}

	return ans
}

func longToIP(x int64, step int) string {
	var ans [4]int
	ans[0], x = int(x&255), x>>8
	ans[1], x = int(x&255), x>>8
	ans[2], x = int(x&255), x>>8
	ans[3] = int(x & 255)
	l := 33
	for step > 0 {
		l--
		step /= 2
	}
	return fmt.Sprintf("%d.%d.%d.%d/%d", ans[3], ans[2], ans[1], ans[0], l)
}

package main

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func runSample(t *testing.T, n int) {

	_, res := solve(n)

	ck := check(n, res)

	if ck != 0 {
		t.Errorf("Sample %d, not correct %d", n, ck)
	}
}

func check(n int, res string) int {
	steps := strings.Split(res, "\n")
	if len(steps) > 20 {
		return 1
	}
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	pattern := regexp.MustCompile(`\d+\+\d+=\d+`)
	num_pattern := regexp.MustCompile(`\d+`)

	var cnt int
	for _, step := range steps {
		insts := pattern.FindAllString(step, -1)
		cnt += len(insts)
		for i := 0; i < len(insts); i++ {
			// a+b=c
			nums := num_pattern.FindAllString(insts[i], -1)
			a, _ := strconv.Atoi(nums[0])
			b, _ := strconv.Atoi(nums[1])
			c, _ := strconv.Atoi(nums[2])

			arr[c-1] = arr[a-1] + arr[b-1]
		}
	}
	var sum int
	for i := 0; i < n; i++ {
		sum += i + 1
		if arr[i] != sum {
			return 2
		}
	}
	if cnt > 2000 {
		return cnt
	}
	return 0
}

func TestSample1(t *testing.T) {
	runSample(t, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 9)
}

func TestSample3(t *testing.T) {
	runSample(t, 100)
}

func TestSample4(t *testing.T) {
	runSample(t, 1000)
}

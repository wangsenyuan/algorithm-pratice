package main

import "testing"

func runSample(t *testing.T, tot int, people [][]int, expect int) {
	score, p := solve(tot, people)

	if score != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, score)
	}

	if expect == 0 {
		return
	}

	k := len(p)
	if k == 0 {
		t.Fatalf("Sample expect %d, but got %v", expect, p)
	}
	var sum int
	var time int
	for _, i := range p {
		if people[i-1][0] >= k {
			sum++
		}
		time += people[i-1][1]
	}

	if sum != expect || time > tot {
		t.Fatalf("Sample result %v, getting wrong result %d %d", p, sum, time)
	}
}

func TestSample1(t *testing.T) {
	tot := 300
	people := [][]int{
		{3, 100},
		{4, 150},
		{4, 80},
		{2, 90},
		{2, 300},
	}
	expect := 2
	runSample(t, tot, people, expect)
}

func TestSample2(t *testing.T) {
	tot := 100
	people := [][]int{
		{1, 787},
		{2, 788},
	}
	expect := 0
	runSample(t, tot, people, expect)
}

func TestSample3(t *testing.T) {
	tot := 100
	people := [][]int{
		{2, 42},
		{2, 58},
	}
	expect := 2
	runSample(t, tot, people, expect)
}

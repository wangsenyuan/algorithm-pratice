package main

import "testing"

func runSample(t *testing.T, people_count, days_in_week, working_day_center, working_day_person int,
	L []int, lunch_begin, lunch_end int, R [][]int, F [][][]int, expect bool) {
	res := solve(people_count, days_in_week, working_day_center, working_day_person, L, lunch_begin, lunch_end, R, F)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	people_count := 2
	days_in_week := 2
	working_day_center := 3
	working_day_person := 2
	L := []int{4, 1}
	begin, end := 2, 3
	R := [][]int{
		{0, 1, 1},
		{0, 1, 0},
	}
	F := [][][]int{
		{
			{1, 1, 1},
			{1, 1, 1},
		},
		{
			{1, 1, 1},
			{1, 0, 1},
		},
	}

	expect := true

	runSample(t, people_count, days_in_week, working_day_center, working_day_person, L, begin, end, R, F, expect)
}

func TestSample2(t *testing.T) {
	people_count := 2
	days_in_week := 2
	working_day_center := 3
	working_day_person := 2
	L := []int{4, 1}
	begin, end := 2, 3
	R := [][]int{
		{0, 1, 2},
		{0, 1, 0},
	}
	F := [][][]int{
		{
			{1, 1, 1},
			{1, 1, 1},
		},
		{
			{1, 1, 1},
			{1, 0, 1},
		},
	}

	expect := false

	runSample(t, people_count, days_in_week, working_day_center, working_day_person, L, begin, end, R, F, expect)
}

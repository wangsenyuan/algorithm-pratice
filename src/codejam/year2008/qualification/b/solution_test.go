package main

import "testing"

func parseTimes(s [][]string) [][]int {
	res := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = make([]int, len(s[i]))
		for j := 0; j < len(s[i]); j++ {
			res[i][j] = ParseTime(s[i][j])
		}
	}
	return res
}

func runSample(t *testing.T, T int, NA, NB int, A [][]string, B [][]string, ea, eb int) {
	AA := parseTimes(A)
	BB := parseTimes(B)

	ra, rb := solve(T, NA, NB, AA, BB)

	if ea != ra || eb != rb {
		t.Errorf("Sample %d %d %d %v %v, expect %d %d, but got %d %d", T, NA, NB, A, B, ea, eb, ra, rb)
	}
}

func TestSample1(t *testing.T) {
	T := 5
	NA, NB := 3, 2
	A := [][]string{
		{"09:00", "12:00"},
		{"10:00", "13:00"},
		{"11:00", "12:30"},
	}
	B := [][]string{
		{"12:02", "15:00"},
		{"09:00", "10:30"},
	}
	runSample(t, T, NA, NB, A, B, 2, 2)
}

func TestSample2(t *testing.T) {
	T := 2
	NA, NB := 2, 0
	A := [][]string{
		{"09:00", "09:01"},
		{"12:00", "12:02"},
	}
	runSample(t, T, NA, NB, A, nil, 2, 0)
}

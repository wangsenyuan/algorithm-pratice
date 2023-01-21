package main

import "testing"

func runSample(t *testing.T, T string, S []string, expect int) {
	res := solve(T, S)

	if (expect < 0) != (len(res) == 0) {
		t.Errorf("Sample expect %d,  but got %v", expect, res)
	}

	if len(res) > 0 {
		if len(res) != expect {
			t.Fatalf("Sample expect %d, but got %v", expect, res)
		}
		color := make([]bool, len(T))
		for _, cur := range res {
			j, w := cur[0], cur[1]
			j--
			w--
			for x := 0; x < len(S[j]); x++ {
				if T[x+w] != S[j][x] {
					t.Fatalf("Sample res %v, not correct", res)
				}
				color[x+w] = true
			}
		}

		for i := 0; i < len(T); i++ {
			if !color[i] {
				t.Fatalf("Sample res %v, not correct", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	T := "bababa"
	S := []string{
		"ba",
		"aba",
	}
	expect := 3
	runSample(t, T, S, expect)
}

func TestSample2(t *testing.T) {
	T := "caba"
	S := []string{
		"bac",
		"acab",
	}
	expect := -1
	runSample(t, T, S, expect)
}

func TestSample3(t *testing.T) {
	T := "abacabaca"
	S := []string{
		"aba",
		"bac",
		"aca",
	}
	expect := 4
	runSample(t, T, S, expect)
}

func TestSample4(t *testing.T) {
	T := "baca"
	S := []string{
		"a",
		"c",
		"b",
	}
	expect := 4
	runSample(t, T, S, expect)
}

func TestSample5(t *testing.T) {
	T := "codeforces"
	S := []string{
		"def",
		"code",
		"efo",
		"forces",
	}
	expect := 2
	runSample(t, T, S, expect)
}

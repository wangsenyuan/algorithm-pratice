package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, posts []string, expect []string) {
	res := solve(A, posts)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2}
	posts := []string{
		"1 1 WhoDoesntLoveChefBook",
		"2 2 WinterIsComing",
		"3 10 TheseViolentDelightsHaveViolentEnds",
		"4 3 ComeAtTheKingBestNotMiss",
	}
	expect := []string{
		"WinterIsComing",
		"WhoDoesntLoveChefBook",
		"TheseViolentDelightsHaveViolentEnds",
		"ComeAtTheKingBestNotMiss",
	}
	runSample(t, A, posts, expect)
}

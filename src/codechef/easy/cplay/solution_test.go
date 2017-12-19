package main

import "testing"

func TestSample1(t *testing.T) {
	s := "10100101111011111111"
	winner, cnt := solve([]byte(s))
	if winner != 1 {
		t.Errorf("sample %s should give winner 1, but got %d", s, winner)
	}
	if cnt != 12 {
		t.Errorf("sample %s should be decided after 12 kicks, but got %d", s, cnt)
	}
}

func TestSample2(t *testing.T) {
	s := "00000000000000000000"
	winner, _ := solve([]byte(s))
	if winner != 0 {
		t.Errorf("sample %s should give winner 0, but got %d", s, winner)
	}
}

func TestSample3(t *testing.T) {
	s := "01011101110110101111"
	winner, cnt := solve([]byte(s))
	if winner != -1 {
		t.Errorf("sample %s should give winner -1, but got %d", s, winner)
	}
	if cnt != 7 {
		t.Errorf("sample %s should be decided after 7 kicks, but got %d", s, cnt)
	}
}

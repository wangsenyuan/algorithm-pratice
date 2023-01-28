package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, items []string, residents []string, expect []string) {
	res := solve(items, residents)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	its := []string{
		"sword weapon 10 2 3 2",
		"pagstarmor armor 0 15 3 1",
		"iceorb orb 3 2 13 2",
		"longbow weapon 9 1 2 1",
	}
	rs := []string{
		"mike gladiator 5 longbow",
		"bobby sentry 6 pagstarmor",
		"petr gladiator 7 iceorb",
		"teddy physician 6 sword",
		"blackjack sentry 8 sword",
		"joe physician 6 iceorb",
	}
	expect := []string{
		"longbow 1 mike",
		"pagstarmor 1 bobby",
		"iceorb 2 joe petr",
	}
	runSample(t, its, rs, expect)
}

func TestSample2(t *testing.T) {
	its := []string{
		"sword weapon 10 2 3 2",
		"pagstarmor armor 0 15 3 1",
		"iceorb orb 3 2 13 2",
		"longbow weapon 9 1 2 1",
	}
	rs := []string{
		"mike gladiator 5 longbow",
		"bobby sentry 6 pagstarmor",
		"petr gladiator 7 iceorb",
		"teddy physician 6 sword",
		"blackjack sentry 8 sword",
	}
	expect := []string{
		"sword 2 mike petr",
		"pagstarmor 1 blackjack",
		"iceorb 2 bobby teddy",
	}
	runSample(t, its, rs, expect)
}

func TestSample3(t *testing.T) {
	its := []string{
		"c armor 0 13 0 3",
		"a weapon 23 0 0 3",
		"b weapon 20 0 0 4",
		"e orb 0 0 13 3",
		"d armor 0 15 0 4",
		"f orb 0 0 17 5",
	}
	rs := []string{
		"j gladiator 7 a",
		"h gladiator 3 f",
		"g gladiator 4 e",
		"i gladiator 7 a",
		"k gladiator 1 b",
	}
	expect := []string{
		"a 3 g i j",
		"d 2 h k",
		"f 0",
	}
	runSample(t, its, rs, expect)
}

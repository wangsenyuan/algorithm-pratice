package p1612

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, name, time []string, expect []string) {
	res := alertNames(name, time)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	name := []string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"}
	time := []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"}
	expect := []string{"daniel"}
	runSample(t, name, time, expect)
}

func TestSample2(t *testing.T) {
	name := []string{"alice", "alice", "alice", "bob", "bob", "bob", "bob"}
	time := []string{"12:01", "12:00", "18:00", "21:00", "21:20", "21:30", "23:00"}
	expect := []string{"bob"}
	runSample(t, name, time, expect)
}

func TestSample3(t *testing.T) {
	name := []string{"john", "john", "john"}
	time := []string{"23:58", "23:59", "00:01"}
	expect := []string{}
	runSample(t, name, time, expect)
}

func TestSample4(t *testing.T) {
	name := []string{"leslie", "leslie", "leslie", "clare", "clare", "clare", "clare"}
	time := []string{"13:00", "13:20", "14:00", "18:00", "18:51", "19:30", "19:49"}
	expect := []string{"clare", "leslie"}
	runSample(t, name, time, expect)
}

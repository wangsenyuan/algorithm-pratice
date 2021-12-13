package p2074

import (
	"reflect"
	"testing"
)

func TestSample1(t *testing.T) {
	robot := Constructor(6, 3)

	robot.Move(2)
	robot.Move(2)

	pos := robot.GetPos()

	if !reflect.DeepEqual(pos, []int{4, 0}) {
		t.Fatalf("Should get {4, 0}, but got %v", pos)
	}

	dir := robot.GetDir()

	if dir != "East" {
		t.Fatalf("Should get East but got %s", dir)
	}

	robot.Move(2)
	robot.Move(1)

	pos = robot.GetPos()

	if !reflect.DeepEqual(pos, []int{5, 2}) {
		t.Fatalf("Should get {5, 2}, but got %v", pos)
	}

	dir = robot.GetDir()

	if dir != "North" {
		t.Fatalf("Should get North but got %s", dir)
	}

	robot.Move(4)

	pos = robot.GetPos()

	if !reflect.DeepEqual(pos, []int{1, 2}) {
		t.Fatalf("Should get {4, 0}, but got %v", pos)
	}

	dir = robot.GetDir()

	if dir != "West" {
		t.Fatalf("Should get West but got %s", dir)
	}

	robot.Move(1)
	pos = robot.GetPos()

	if !reflect.DeepEqual(pos, []int{0, 2}) {
		t.Fatalf("Should get {0, 2}, but got %v", pos)
	}

	dir = robot.GetDir()

	if dir != "West" {
		t.Fatalf("Should get West but got %s", dir)
	}

	robot.Move(2)
	pos = robot.GetPos()

	if !reflect.DeepEqual(pos, []int{0, 0}) {
		t.Fatalf("Should get {0, 0}, but got %v", pos)
	}

	dir = robot.GetDir()

	if dir != "South" {
		t.Fatalf("Should get South but got %s", dir)
	}

}

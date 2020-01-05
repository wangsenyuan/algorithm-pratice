package p1311

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, watchedVideos [][]string, friends [][]int, id int, level int, expect []string) {
	res := watchedVideosByFriends(watchedVideos, friends, id, level)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	watchedVideos := [][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}}
	friends := [][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}}
	id := 0
	level := 1
	expect := []string{"B", "C"}
	runSample(t, watchedVideos, friends, id, level, expect)
}

func TestSample2(t *testing.T) {
	watchedVideos := [][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}}
	friends := [][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}}
	id := 0
	level := 2
	expect := []string{"D"}
	runSample(t, watchedVideos, friends, id, level, expect)
}

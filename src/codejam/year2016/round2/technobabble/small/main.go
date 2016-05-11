package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type topic struct {
	first, second string
}

var fx = 0
var sx = 0

func play(topics []topic) int {
	fs := make(map[string]struct{})
	ss := make(map[string]struct{})

	for _, t := range topics {
		fs[t.first] = struct{}{}
		ss[t.second] = struct{}{}
	}
	fx = len(fs)
	sx = len(ss)

	n := len(topics)
	unFaked, _ := visit(make([]topic, 0, n), topics)

	return n - len(unFaked)
}

func copyMap(s map[string]struct{}) map[string]struct{} {
	t := make(map[string]struct{})
	for k, v := range s {
		t[k] = v
	}
	return t
}

func visit(result []topic, topics []topic) ([]topic, bool) {
	fs := make(map[string]struct{})
	ss := make(map[string]struct{})
	for _, t := range result {
		fs[t.first] = struct{}{}
		ss[t.second] = struct{}{}
	}

	if len(fs) == fx && len(ss) == sx {
		return result, true
	}

	if len(topics) == 0 {
		return nil, false
	}

	result1, found1 := visit(append(result, topics[0]), topics[1:])
	result2, found2 := visit(result, topics[1:])

	if !found1 && !found2 {
		return nil, false
	}

	if !found1 {
		return result2, true
	}

	if !found2 {
		return result1, true
	}

	if len(result1) < len(result2) {
		return result1, true
	}

	return result2, true
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	t, _ := strconv.Atoi(strings.Trim(line, "\n "))

	for i := 1; i <= t; i++ {
		line, _ = reader.ReadString('\n')
		n, _ := strconv.Atoi(strings.Trim(line, "\n"))
		topics := make([]topic, 0, n)
		for j := 0; j < n; j++ {
			line, _ = reader.ReadString('\n')
			words := strings.Split(line, " ")
			topic := topic{words[0], words[1]}
			topics = append(topics, topic)
		}
		r := play(topics)
		fmt.Printf("Case #%d: %d\n", i, r)
	}
}

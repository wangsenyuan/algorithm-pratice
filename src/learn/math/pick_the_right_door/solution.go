package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	strategy := flag.String("s", "switch", "switch the door or not, when given the hint")
	round := flag.Int("c", 10000, "number of rounds to simulate")

	flag.Parse()

	simulate(*strategy == "switch", *round)
}

func simulate(change bool, round int) {
	var count int

	for i := 0; i < round; i++ {
		winDoor := randomChoose()
		pick := randomChoose()
		res := pick
		if change {
			res = changeTheDoor(winDoor, pick)
		}
		fmt.Fprintf(os.Stderr, "round %d: %d %d %d\n", i, winDoor, pick, res)
		if res == winDoor {
			count++
		}
	}
	x := float64(count) / float64(round) * 100
	fmt.Printf("X = %.3f\n", x)
	o := 100.0 - x
	fmt.Printf("O = %.3f\n", o)
}

func randomChoose() int {
	num := rand.Intn(90)

	return num % 3
}

func changeTheDoor(win, pick int) int {
	doors := make([]int, 0, 2)
	for i := 0; i < 3; i++ {
		if i == pick {
			continue
		}
		doors = append(doors, i)
	}

	if doors[0] == win {
		return doors[0]
	}

	if doors[1] == win {
		return doors[1]
	}

	num := rand.Intn(100)
	j := num % 2
	return doors[j]
}

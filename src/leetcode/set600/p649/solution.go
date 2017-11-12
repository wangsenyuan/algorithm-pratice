package main

import "fmt"

func predictPartyVictory(senate string) string {
	ban := make([]bool, len(senate))

	for {
		sameParty := true

		for i := 1; i < len(senate); i++ {
			if senate[i] != senate[0] {
				sameParty = false
				break
			}
		}

		if sameParty {
			break
		}

		for i := 0; i < len(senate); i++ {
			ban[i] = false
		}

		j := 1
		for i := 0; i < len(senate); i++ {
			if ban[i] {
				continue
			}

			for j < 2*len(senate) && senate[j%len(senate)] == senate[i] || ban[j%len(senate)] {
				j++
			}

			if j%len(senate) == i {
				break
			}
			ban[j%len(senate)] = true
			j++
		}

		var tmp string
		for i := 0; i < len(senate); i++ {
			if !ban[i] {
				tmp += string(senate[i])
			}
		}
		senate = tmp
		//fmt.Println(senate)
	}

	if senate[0] == 'R' {
		return "Radiant"
	}
	return "Dire"
}

func main() {
	fmt.Println(predictPartyVictory("RD"))
	//fmt.Println(predictPartyVictory("RDD"))
	//fmt.Println(predictPartyVictory("DDRRRR"))
	//fmt.Println(predictPartyVictory("DRDRR"))
	fmt.Println(predictPartyVictory("DRRDRDRDRDDRDRDRD"))
}

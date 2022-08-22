package p2383

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {

	var day int

	n := len(energy)

	cur_energy := initialEnergy
	cur_exp := initialExperience

	for i := 0; i < n; i++ {
		if cur_energy <= energy[i] {
			day += energy[i] - cur_energy + 1
			cur_energy = energy[i] + 1
		}
		if cur_exp <= experience[i] {
			day += experience[i] - cur_exp + 1
			cur_exp = experience[i] + 1
		}
		cur_energy -= energy[i]
		cur_exp += experience[i]
	}

	return day
}

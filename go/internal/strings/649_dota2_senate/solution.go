package dota2_senate

func predictPartyVictory(senate string) string {
	radiant := make([]int, 0, len(senate))
	dire := make([]int, 0, len(senate))
	for i, c := range senate {
		if c == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}

	for len(radiant) > 0 && len(dire) > 0 {
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+len(senate))
		} else {
			dire = append(dire, dire[0]+len(senate))
		}
		radiant = radiant[1:]
		dire = dire[1:]
	}

	if len(radiant) > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}

var Solution = predictPartyVictory

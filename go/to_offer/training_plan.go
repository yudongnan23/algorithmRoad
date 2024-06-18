package to_offer

func trainingPlan(actions []int) int {
	m := make(map[int]int, 0)
	for i := 0; i < len(actions); i++ {
		if m[actions[i]] == 2 {
			delete(m, actions[i])
			continue
		}
		m[actions[i]] = m[actions[i]] + 1
	}

	for k, _ := range m {
		return k
	}
	return 0
}

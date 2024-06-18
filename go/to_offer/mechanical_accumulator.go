package to_offer

func mechanicalAccumulator(target int) int {
	_ = target > 0 && func() bool {
		target = target + mechanicalAccumulator(target-1)
		return true
	}()
	return target
}

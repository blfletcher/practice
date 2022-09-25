package main

// https://leetcode.com/problems/minimum-health-to-beat-game/

func minimumHealth(damage []int, armor int) int64 {
	max := 0
	damageSum := int64(0)

	for i := range damage {
		dm := damage[i]

		if dm > max {
			max = dm
		}

		damageSum += int64(dm)
	}

	arm := armor
	if max < armor {
		arm = max
	}

	return damageSum - int64(arm) + 1
}

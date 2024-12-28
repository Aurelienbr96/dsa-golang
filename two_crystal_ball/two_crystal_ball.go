package twocrystalball

import (
	"math"
)

func TwoCrystalBall(arr []bool) int {
	jump_val := int(math.Floor(math.Sqrt(float64(len(arr)))))

	for i := 1; i < jump_val+1; i++ {
		prevIndex := jump_val * (i - 1)

		if !arr[jump_val*i] {
			continue
		} else {
			for a := 0; a < jump_val; a++ {
				if arr[prevIndex+a] {
					return jump_val*(i-1) + a
				}
			}
		}
	}

	return jump_val
}

package easing

func T(k, min, max float64) float64 {
	if max <= min {
		return 0
	}
	return (k - min) / (max - min)
}

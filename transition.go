package easing

import (
	"math"
)

func Transition(from, to, t float64) float64 {
	return from + (to-from)*t
}

func TransitionUint(from, to uint, t float64) uint {
	x := Transition(float64(from), float64(to), t)
	return uint(math.Floor(x))
}

func TransitionInt(from, to int, t float64) int {
	x := Transition(float64(from), float64(to), t)
	return int(math.Floor(x))
}

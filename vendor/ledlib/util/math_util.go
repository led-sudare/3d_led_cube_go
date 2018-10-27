package util

import "math"

func RoundToInt(input float64) int {
	return int(input + 0.5)
}

func FloorToInt(input float64) int {
	return int(math.Floor(input))
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func GetSign(value float64) int {
	if value >= 0 {
		return 1
	} else {
		return -1
	}
}

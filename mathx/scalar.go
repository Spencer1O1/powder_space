package mathx

import "math"

func Sqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}

func Min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func Max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func FastFloorToInt(x float32) int {
	i := int(x)
	if float32(i) > x {
		return i - 1
	}
	return i
}

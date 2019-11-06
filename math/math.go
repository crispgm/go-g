package math

import "math"

// Deg2Rad convert degree to radian
func Deg2Rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

// Rad2Deg convert radian to degree
func Rad2Deg(deg float64) float64 {
	return deg * 180 / math.Pi
}

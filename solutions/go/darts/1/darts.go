package darts

import "math"

// Score returns the points earned by a dart landing at coordinate (x, y).
func Score(x, y float64) int {
	d := math.Hypot(x, y)
	switch {
	case d <= 1:
		return 10
	case d <= 5:
		return 5
	case d <= 10:
		return 1
	default:
		return 0
	}
}


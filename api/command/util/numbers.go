package util

// MaxInt は大きい方のintを返す。
func MaxInt(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

// MinInt は小さい方のintを返す。
func MinInt(v1, v2 int) int {
	if v1 > v2 {
		return v2
	}
	return v1
}

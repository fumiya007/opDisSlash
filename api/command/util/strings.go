package util

import "strings"

// SplitMultiSep は複数のセパレータのスプリットを行う
func SplitMultiSep(s string, seps ...string) []string {
	if len(seps) == 0 {
		return []string{s}
	}

	sep0 := seps[0]
	for i := 1; i < len(seps); i++ {
		s = strings.ReplaceAll(s, seps[i], sep0)
	}

	return strings.Split(s, sep0)
}

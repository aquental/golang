package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
    result := make(map[string]int)
	for score, vals := range in {
		for _, val := range vals {
			result[strings.ToLower(val)] = score
		}
	}
	return result
}


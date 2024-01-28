package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var list []string
    for _, item := range candidates {
		if SortStringIgnoreCase(subject) == SortStringIgnoreCase(item) && !strings.EqualFold(subject,item) {
			list = append(list, item)
		}
    }
    return list
}

func SortStringIgnoreCase(w string) string {
    s := []rune(strings.ToLower(w))
    sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })

    return string(s)
}

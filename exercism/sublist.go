package sublist
import (
    "strconv"
    "strings"
)
// Relation type is defined in relations.go file.
// Sublist returns the relation between two lists l1 and l2.
func Sublist(l1, l2 []int) Relation {
	// Return RelationEqual if both lists are empty
	if len(l1) == 0 && len(l2) == 0 {
		return RelationEqual
	}
	// Return RelationSuperlist if l1 is not empty and l2 is empty
	if len(l1) > 0 && len(l2) == 0 {
		return RelationSuperlist
	}
	// Return RelationSublist if l1 is empty and l2 is not empty
	if len(l1) == 0 && len(l2) > 0 {
		return RelationSublist
	}

	// Convert list elements to strings
	var s1, s2 string
	for _, n := range l1 {
		s1 += strconv.Itoa(n) + ","
	}
	for _, n := range l2 {
		s2 += strconv.Itoa(n) + ","
	}

	// Return RelationEqual if the string representations of l1 and l2 are equal
	if s1 == s2 {
		return RelationEqual
	}
	// Return RelationSuperlist if s1 contains s2
	if strings.Contains(s1, s2) {
		return RelationSuperlist
	}
	// Return RelationSublist if s2 contains s1
	if strings.Contains(s2, s1) {
		return RelationSublist
	}
	// Return RelationUnequal if none of the above conditions are met
	return RelationUnequal
}

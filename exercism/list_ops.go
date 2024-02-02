package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

// Foldl applies a function to each element of the IntList in a left-to-right order
func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	// Iterate through the elements of the IntList
	for _, v := range s {
		// Apply the function to the current element and the accumulator
		initial = fn(initial, v)
	}
	// Return the final result
	return initial
}

// Foldr applies the given function fn to each element of the IntList in reverse order
// starting with the initial value, and returns the accumulated result.
func (s IntList) Foldr(fn func(int, int) int, initial int) int {
    for i := len(s) - 1; i >= 0; i-- {
        initial = fn(s[i], initial)
    }
    return initial
}

// Filter applies the given function to each element in the list and returns a new list 
// containing only the elements for which the function returns true.
func (s IntList) Filter(fn func(int) bool) IntList {
    // Initialize an empty list to store the filtered elements
    res := IntList{}
    // Iterate through the elements in the original list
    for _, v := range s {
        // Apply the function to the current element and append it to the result list if it returns true
        if fn(v) {
            res = append(res, v)
        }
    }
    // Return the filtered list
    return res
}

// Length returns the number of elements in the IntList.
func (s IntList) Length() int {
	return len(s)
}

// Map applies the given function to each element of the IntList and returns a new IntList.
func (s IntList) Map(fn func(int) int) IntList {
    result := make(IntList, len(s))
    for i, v := range s {
        result[i] = fn(v)
    }
    return result
}

// Reverse reverses the order of elements in the IntList.
func (s IntList) Reverse() IntList {
	// Initialize left and right pointers
	l := 0
	r := len(s) - 1

	// Swap elements until the pointers meet
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}

	// Return the reversed IntList
	return s
}

// Append appends the elements of lst to the IntList s.
func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

// Concat concatenates the elements of the input IntList with the current IntList
func (s IntList) Concat(lists []IntList) IntList {
	// iterate through the input lists
	for _, l := range lists {
		// append the elements of each list to the current IntList
		s = s.Append(l)
	}
	// return the concatenated IntList
	return s
}

package proverb

// Proverb will receive a list of words and create a proverb.
func Proverb(rhyme []string) []string {
    length := len(rhyme)
    if length < 1 {
        return nil
    }
    
	var res []string
	for i, val := range rhyme {
		if i < length-1 {
			res = append(res, "For want of a "+val+" the "+string(rhyme[i+1])+" was lost.")
		} else {
			res = append(res, "And all for the want of a "+string(rhyme[0])+".")
		}
	}
	return res
}


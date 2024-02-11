package wordy

import (
	"fmt"
	"strings"
    "strconv"
)

const (
    by string = "by"
    plus string = "plus"
    minus string = "minus"
    multiplied string = "multiplied"
    divided string = "divided"
    
)
func Answer(question string) (int, bool) {
    var stack[]int

    question = strings.Trim(strings.ToLower(question),"!?")
    fmt.Printf(": %s\n",question)
	for _, w := range strings.Split(question, " ")	{
		fmt.Println(w)
        i, err := strconv.Atoi(w)
        if err == nil {
            stack = append(stack,i)
            fmt.Printf("-> %v\n",stack)
        } else {
        	switch w {
                case by:
            		fmt.Println("=> "+w)
    			case plus:
        			fmt.Println(" > "+w)
    			case minus:
        			fmt.Println(" > "+w)
    			case multiplied:
        			fmt.Println(" > "+w)
                case divided:
        			fmt.Println(" > "+w)
    		}
        }
        
	}
	return 5, true
}

package wordy

import (
	"fmt"
	"strconv"
	"strings"
)

const (
    by string = "by"
    plus string = "plus"
    minus string = "minus"
    multiplied string = "multiplied"
    divided string = "divided"
    
)

func fPlus(stack[]int) int {
	//pop two items from the stack and push the result
    len := len(stack)
    fmt.Printf("fPlus: len %d - %v\n",len,stack)
    if len < 2 {
        return 0;
    }
    p1 := stack[len-1]
    p2 := stack[len-2]
    result := p1 + p2
    fmt.Printf("fPlus: p1: %d, p2: %d, result: %d\n",p1,p2,result)

    return result
}

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

package wordy

import (
	"fmt"
	"strconv"
	"strings"
)

type stack []int
type opsFunc func(stack) bool

func (s stack) Push(v int) stack {
    var tmp stack
    tmp = append(s, v)
    fmt.Printf("Push(%d) -> %v\n", v, tmp)
    return tmp
}

func (s stack) Pop() (stack, int) {
    fmt.Printf("Pop() -> ")
    l := len(s)
    fmt.Printf("len=%d, ",l)
    fmt.Printf("%v, %d\n", s, s[l-1])
    return  s[:l-1], s[l-1]
}

func (s stack) Len() (int) {
    l := len(s)
    fmt.Printf("Len() -> %d\n",l)
    return l
}

func (s stack) Calculate (fun func(int, int)int) (stack,bool) {
    //ans := fp(3,2)
    fmt.Printf("Calculate() -> %v\n",s)
    s, p1 := s.Pop()
    s, p2 := s.Pop()
    //push result
    s = s.Push(fun(p1,p2))
    fmt.Printf("Calculate: p1: %d, p2: %d, result: %d\n",p1,p2,fun(p1,p2))
    //s.Print()

    return s,true
}

func (s stack) Print() {
    fmt.Printf("Print() -> %v\n",s)
}

const (
    by string = "by"
    plus string = "plus"
    minus string = "minus"
    multiplied string = "multiplied"
    divided string = "divided"
    cubed string = "cubed"
    who string = "who"
)

func fPlus(p2, p1 int) int {
    return (p1 + p2)
}

func fMinus(p2, p1 int) int {
    return (p1 - p2)
}

func fMult(p2, p1 int) int {
    return (p1 * p2)
}

func fDiv(p2, p1 int) int {
    return (p1 / p2)
}

func Answer(question string) (int, bool) {
    var s stack
    var f func(int, int)int = nil

    question = strings.Trim(strings.ToLower(question),"!?")
    fmt.Printf(": %s\n",question)
	for _, w := range strings.Split(question, " ")	{
		fmt.Println(w)
        i, err := strconv.Atoi(w)
        if err == nil {
            s = s.Push(i)
            //fmt.Printf("-> %v\n",s)
            if f != nil {
                var res bool
                s, res = s.Calculate(f)
                f = nil
                if ! res {
                    return 0, false
                }
            }
        } else {
        	switch w {
    			case plus:
                	f = fPlus
                    fmt.Printf(" > %s, %p",w, f)
    			case minus:
                	f = fMinus
                	fmt.Printf(" > %s, %p",w, f)
    			case multiplied:
                	f = fMult
                	fmt.Printf(" > %s, %p",w, f)
                case divided:
                	f = fDiv
                	fmt.Printf(" > %s, %p",w, f)
            	case cubed:
					return 0, false
                case "who":
            		return 0, false
    		}
        }
	}
	
	s, ret := s.Pop()
    if ! ret {
        return 0, false
    }
	return ret, true
}

package main2

import (
	"fmt"
	"time"
)
func say(s string) {
	for i:=0; i< 5; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
	}
}

func main() {
	go say("-> hello")
	say("world")
}

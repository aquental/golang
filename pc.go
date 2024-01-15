package main

import (
	"fmt"
	"math"
	"math/rand"
)

type primemsg struct {
	num int
	isPrime bool
}

const TEST_LEN = 100
const WRKS = 4

func isPrime(cin chan primemsg, cout chan primemsg, id int) {
	i:=0
	
	for {
		msg := <- cin
		num := msg.num
		fmt.Println(id, "] is checking ", num)
		sq_root := int(math.Sqrt(float64(num)))
		for i = 2; i <= sq_root; i++ {
			if num % i == 0 {
				msg.isPrime = false
				cout <- msg
				break
			}
		}
		if i > sq_root {
			msg.isPrime = true
			cout <- msg
		}
	}
}

func main(){
	msg := primemsg{num: 0, isPrime: false}

	//create channels
	cin := make(chan primemsg)
	cout := make(chan primemsg)

	//create workers
	for i := 0; i < WRKS; i++ {
		go isPrime(cin, cout, i+1)
	}

	//fill in
	for i := 0; i < TEST_LEN; i++ {
		msg.num = rand.Intn(1000000)+1000000
		cin <- msg
	}

	//get results out
	for i := 0; i < TEST_LEN; i++ {
		msg = <- cout
		fmt.Println(msg.num, msg.isPrime)
	}
}

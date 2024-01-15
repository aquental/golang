package main1

import (
	"fmt"
)

func sum (c chan int){
	x := <- c
	y := <- c
	c <- x + y
}

func sum2(c chan int, o chan int){
	for  {
		x:= <- c
		y:= <- c
		fmt.Println(x, "+",y,"=",x+y)
		o <- x+y
	}
}

func main (){

	c := make(chan int, 2)
	go sum(c)
	c <- 10
	c <- 15
	r := <- c
	fmt.Println(r)

	fmt.Println("---")

	i := make(chan int, 6)
	o := make(chan int, 3)
	go sum2(i,o)
	i <- 10
	i <- 15
	i <- 99
	i <- 1
	i <- 23
	i <- 7

	r = <- o
	fmt.Println(r)
	r = <- o
	fmt.Println(r)
	r = <- o
	fmt.Println(r)
}

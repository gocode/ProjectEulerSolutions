package main

import "fmt"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	r1 := make(chan int)
	r2 := make(chan int)
	go sumOfSquare(c1, r1)
	go squareOfSum(c2, r2)
	for i := 1; i <= 100; i++ {
		c1 <- i
		c2 <- i
	}
	close(c1)
	close(c2)
	d := <-r1 - <-r2
	fmt.Printf("Difference between the sum of the squares of the first one hundred natural numbers and the square of the sum %d", d)
}

func sumOfSquare(c, r chan int) {
	s := 0
	for i := range c {
		s = s + i*i
	}
	r <- s
}

func squareOfSum(c, r chan int) {
	s := 0
	for i := range c {
		s = s + i
	}
	r <- s * s
}

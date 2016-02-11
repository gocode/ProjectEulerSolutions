package main

import "fmt"

const N int64 = 600851475143

func main() {
	//setup prime series
	c := make(chan int64)
	go PrimeSeries(c)

	var n int64 = 600851475143
	for {
		p := <-c
		if n%p == 0 {
			n = n / p
			if isPrime(n) {
				if p > n {
					n = p
				}
				fmt.Printf("The largest prime factor is %d", n)
				return
			}
		}
	}
}

func PrimeSeries(c chan<- int64) {
	var i int64
	for i = 2; i < N/2; i++ {
		if isPrime(i) {
			c <- i
		}
	}

}

func isPrime(n int64) bool {
	var i int64
	if n == 2 {
		return true
	}
	for i = 2; i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

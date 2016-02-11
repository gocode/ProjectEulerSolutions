package main

import "fmt"

func main() {
	pn := 0
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			n := i * j
			if isPallandrome(n) && n > pn {
				pn = n
			}
		}
	}
	fmt.Printf("Largest product of two three-digit numbers %d", pn)
}

func isPallandrome(n int) bool {
	s := 0
	tn := n
	for {
		a := n % 10
		n = n / 10
		s = s*10 + a
		if n <= 0 {
			break
		}
	}
	if s == tn {
		return true
	} else {
		return false
	}
}

package main

import "fmt"

func main() {
	prime := []int{2, 3, 5, 7, 11, 13, 17, 19}
	nonprime := []int{4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20}
	isDivisible := true
	isNonPrimeDivisible := true
	for i := 40; ; i++ {
		for j := 0; j < len(prime); j++ {
			if i%prime[j] != 0 {
				isDivisible = false
				break
			}
		}
		if isDivisible {
			for k := 0; k < len(nonprime); k++ {
				if i%nonprime[k] != 0 {
					isNonPrimeDivisible = false
					break
				}
			}
			if isNonPrimeDivisible {
				fmt.Printf("The smallest number that can be divided from 0 - 20 is %d", i)
			}
		}
		isDivisible = true
		isNonPrimeDivisible = true
	}
}

// https://projecteuler.net/problem=5
//
// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

import "fmt"

func prime_factors(x int, limit int) bool {
	

}

func solve(limit int) int {
	
	for i := 1; true; i++ {
		if divisible(i, limit){
			return i
		}

	}
	return 0
}

func main() {
	fmt.Printf("%d", solve(20))

}

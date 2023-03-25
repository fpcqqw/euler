package main

import (
	"fmt"
	"math"
)

func prime_factors() {


}

func prime_sieve_lookup(limit int) []bool {
	prime := make([]bool, limit)
	for i := 1; i < limit; i++ {
		prime[i] = true
	}
	prime[0] = false
	
	for p:=2; math.Pow(float64(p),2)<float64(limit);p++{
		fmt.Printf("p: %d", p)
		if prime[p-1] {
			for i:=int(math.Pow(float64(p),2));i<=limit;i+=p{
				fmt.Printf("i: %d, %d", i, p)
				prime[i-1] = false
			}
		} 
	} 
	return prime
}

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

func prieme_sieve(limit int) []int{
	lookup := prime_sieve_lookup(limit)
	fmt.Printf("%v", lookup)
	result := []int{} 
	for i, is_prime := range lookup {
		if is_prime{
			result = append(result, i+1)
		}

	}
	return result
}

func main() {
	fmt.Printf("%v", prieme_sieve(100000))
}

// https://projecteuler.net/problem=3

// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

package main

import "fmt"

func calc_prime_factors(num int) []int {
	prime_nums := prime_sieve(num)
	prime_factors := []int{}

	for _, prime_num:= range prime_nums {
		if num % prime_num == 0{
			prime_factors = append(prime_factors, prime_num)
			num = num/prime_num
		}
	}
	return prime_factors
}

func main() {
	fmt.Printf("hi")
}
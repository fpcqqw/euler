package main

import (
	"fmt"
	"math"
)

func prime_sieve_lookup(limit int) []bool {
	prime := make([]bool, limit)
	for i := 1; i < limit; i++ {
		prime[i] = true
	}
	fmt.Printf("halooooooooo")
	prime[0] = false

	for p := 2; math.Pow(float64(p), 2) < float64(limit); p++ {
		//fmt.Printf("p: %d\n", p)
		if prime[p-1] {
			for i := int(math.Pow(float64(p), 2)); i <= limit; i += p {
				//fmt.Printf("i: %d, %d\n", i, p)
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

func prime_sieve(limit int) []int {
	lookup := prime_sieve_lookup(limit)
	//fmt.Printf("%v", lookup)
	result := []int{}
	for i, is_prime := range lookup {
		if is_prime {
			result = append(result, i+1)
		}

	}
	return result
}


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


// func prime_factors(num int) []int {
// 	prime_nums := prime_sieve(num)
// 	// divide num by smallest prime num
// 	// if possible, add to list of prime factors
// 	prime_factors := []int{}
// 	if num != 0{
// 		num, factor = get_prime_factor(num, prime_nums)
// 		prime_factors = append(prime_factors, factor)
// 	}
// 	return prime_nums
// }

// func get_prime_factor(num int, prime_nums []int) []int{
// 	prime_factors := []int{}

// 	for p := range prime_nums {
// 		fmt.Printf("Current num: %d \n", num)
// 		if num % prime_nums[p] == 0{
// 			fmt.Printf("Prime factor found: %d \n", prime_nums[p])
// 			prime_factors = append(prime_factors, prime_nums[p])
// 			num = num / prime_nums[p]
// 		} 

// 	}
// 	return prime_factors
// }

func main() {
	fmt.Printf("%v", calc_prime_factors(600851475143/2))
}
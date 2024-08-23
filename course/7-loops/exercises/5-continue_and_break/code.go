package main

import (
	"fmt"
  "math"
)

func printPrimes(max int) {
  for i := 2; i <= max; i++ {
    if i == 2{
      fmt.Println(i)
      continue
    }
    if i & 1 == 0 {
      continue
    }
    prime := true
    for n := 3; n <= int(math.Pow(float64(i), .5 )); n++ {
      if i % n == 0 {
        prime = false
        break
      }
    }
    if prime {
      fmt.Println(i)
    }
  }

}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}

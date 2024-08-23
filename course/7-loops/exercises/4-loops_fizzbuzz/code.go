package main

import "fmt"

func fizzbuzz() {
  for i:=1; i<=100; i++{
    if i % 3 != 0 && i %5 != 0{
      fmt.Println(i)
    }else{
      if i % 3 == 0 {
        fmt.Print("fizz")
      }
      if i % 5 == 0 {
        fmt.Print("buzz")
      }
      fmt.Print("\n")
    }
  }
}

// don't touch below this line

func main() {
	fizzbuzz()
}

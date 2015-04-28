package main

import (
  "fmt"
  "os"
  "strconv"
)

func generate(limit int) <-chan int {
  c := make(chan int)
  
  go func() {
    c <- 2
    for i := 3; i <= limit; i += 2 {
      c <- i
    }
    close(c)
  }()
  
  return c
}

func primes(g <-chan int) <-chan int {
  p := make(chan int)
  
  go func() {
    for n := range g {
      check(n, p)
    }
    close(p)
  }()
  
  return p
}

func check(n int, p chan int){
  for i := 2; i <= n / 2; i++ {
    if n % i == 0 {
      return
    }
  }
  p <- n
}

func output(p int){
  fmt.Println(p)
}

func main() {
  limit, _ := strconv.Atoi(os.Args[1])

  for p := range primes(generate(limit)) {
    output(p)
  }
}

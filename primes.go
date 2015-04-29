package main

import (
  "fmt"
  "os"
  "runtime"
  "strconv"
)

func generate(limit int, g chan<- int, q chan<- bool) {
  g <- 2
  for i := 3; i <= limit; i += 2 {
    g <- i
  }
  q <- true
}

func primes(g <-chan int, p chan<- int) {
  for n := range g {
    check(n, p)
  }
}

func check(n int, p chan<- int){
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
  runtime.GOMAXPROCS(runtime.NumCPU())

  limit, _ := strconv.Atoi(os.Args[1])
  g := make(chan int)
  p := make(chan int)
  q := make(chan bool)

  go generate(limit, g, q)
  for i := 0; i < runtime.NumCPU(); i++ {
    go primes(g, p)
  }

  go func() {
    if <-q {
      close(g)
      close(p)
    }
  }()

  for n := range p {
    output(n)
  }
}

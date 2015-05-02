package main

import (
  "fmt"
  "math"
  "os"
  "reflect"
  "runtime"
  "strconv"
)

func generate(limit int, g chan<- int) {
  g <- 2
  for i := 3; i <= limit; i += 2 {
    g <- i
  }
  close(g)
}

func primes(g <-chan int, p chan<- int) {
  for n := range g {
    check(n, p)
  }
  close(p)
}

func check(n int, p chan<- int){
  max := int(math.Sqrt(float64(n)))
  for i := 3; i <= max; i++ {
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
  numCPUs := runtime.NumCPU()
  runtime.GOMAXPROCS(numCPUs)

  limit, _ := strconv.Atoi(os.Args[1])

  g := make(chan int, limit)
  p := make([]chan int, numCPUs)
  waiting := make(map[int]bool)

  go generate(limit, g)
  for i := 0; i < numCPUs; i++ {
    p[i] = make(chan int)
    go primes(g, p[i])
    waiting[i] = true
  }

  cases := make([]reflect.SelectCase, len(p))
  for i, ch := range p {
    cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
  }

  for len(waiting) > 0 {
    i, n, ok := reflect.Select(cases)

    if ok {
      output(int(n.Int()))
    } else {
      delete(waiting, i)
    }
  }
}

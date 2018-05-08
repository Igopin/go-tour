package main

import "fmt"

func fibonacci2() func() int {
  x, y := 0, 1
  
  return func() int {
      res := x
      x, y = y, x + y
      return res
  }
}

// my first version
func fibonacci() func() int {
  calls, sum, p1, p2 := 0, 0, 1, 0
  
  return func() int {
    switch calls {
    case 0:
        calls++
        sum = 0
    case 1:
        calls++
        sum = 1
    default:
       sum = p1 + p2
       p2, p1 = p1, sum
    }
    return sum
  }
}

func main() {
  f := fibonacci2()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}

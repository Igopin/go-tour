package main

import (
    "fmt"
    "time"
)

func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        fmt.Print('.')
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println("wait")
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    
    time.Sleep(5000 * time.Millisecond)
    fibonacci(c, quit)
}

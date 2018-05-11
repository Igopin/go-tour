package main

import (
    "golang.org/x/tour/tree"
    "fmt"
)

func WalkProc(t *tree.Tree, ch chan int) {
    if t == nil {
        return
    }
    WalkProc(t.Left, ch)
    ch <- t.Value
    WalkProc(t.Right, ch)
}

func Walk(t *tree.Tree, ch chan int) {
   WalkProc(t, ch)
   close(ch)
}

func Same(t1, t2 *tree.Tree) (res bool) {
    res = true

    ch1, ch2 := make(chan int), make(chan int)
    go Walk(t1, ch1)
    go Walk(t2, ch2)

    v1, ok1 := <-ch1
    v2, ok2 := <-ch2

    for ok1 && ok2 {
        if v1 != v2 {
            res = false
            break
        }
        v1, ok1 = <-ch1
        v2, ok2 = <-ch2
    }

    if ok1 || ok2 {
        res = false
    }

    return res
}

func main() {
    ch := make(chan int)
    go Walk(tree.New(1), ch)

    for i := 0; i < 10; i++ {
        fmt.Println(<-ch)
    }

    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
}

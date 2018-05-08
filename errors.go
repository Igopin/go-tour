package main

import (
    "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt)Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
    var err error
    var z float64

    if x < 0 {
        err = ErrNegativeSqrt(x)
    } else {
        z = 1.0
        for i := 0; i < 10; i++ {
            z -= (z*z - x) / (2*z)
        }
    }
    return z, err
}

func main() {
    if res, err := Sqrt(-2); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(res)
    }

    fmt.Println(Sqrt(2))
}

package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

const Shift byte = 13

func ShiftLetter( char byte, shift byte ) byte {
    var res byte
    switch {
    case 'A' <= char && char <= 'Z':
        if char + shift > 'Z' {
            res = 'A' + (char + shift  - 'A') % 26
        } else {
            res = char + shift
        }
    case 'a' <= char && char <= 'z':
        if char + shift > 'z' {
            res = 'a' + (char + shift - 'a') % 26
        } else {
            res = char + shift
        }
    default:
        res = char
    } 

    return res
}


func (r13 rot13Reader)Read(b []byte) (int, error) {
    if len(b) == 0 {
        return 0, nil
    }

    num, err := r13.r.Read(b)
    
    for i := 0; i < num; i++ {
        b[i] = ShiftLetter(b[i], Shift)
    }
    return num, err
}


func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}

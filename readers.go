package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader)Read(a []byte)(int, error) {
    var num int = 0
    var err error

    if len(a) == 0 {
        return 0, nil
    }

    for i, _ := range a {
        a[i] = 'A'
        num++
    }

    return num, err
}

func main() {
    reader.Validate(MyReader{})
}

package main

import (
    "strings"
    "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
    chars := make(map[string]int)

    for _, word := range strings.Fields(s) {
        chars[word]++
    }
    return chars
}

func main() {
    wc.Test(WordCount)
}

package main

import (
    "fmt"
    "time"
    "sync"
)

type Visited struct {
    urls map[string]struct{}
    mux sync.Mutex
} 

var visited = Visited{urls: make(map[string]struct{})}

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
    if depth <= 0 {
        return
    }

    visited.mux.Lock()
    if _, ok := visited.urls[url]; ok {
      visited.mux.Unlock()
      return
    }

    visited.urls[url] = struct{}{}
    visited.mux.Unlock()

    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    fmt.Printf("found: %s %q\n", url, body)
    fetched := make(chan bool)
    for i, u := range urls {
        fmt.Printf("-> Start fetch child %v/%v: %v\n", i, len(urls), urls[i])
        go func(url string) {
            Crawl(url, depth-1, fetcher)
            fetched <- true
        }(u)
    }
    for _, u := range urls {
        <- fetched
        fmt.Printf("<- [%v]: %v\n", url, u)
    }
    //return
}

func main() {
    Crawl("https://golang.org/", 4, fetcher)
    fmt.Println("Result:")
    for key, _ := range visited.urls {
        fmt.Println(key)
    }
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
    time.Sleep(500 * time.Millisecond)
    if res, ok := f[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
    "https://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "https://golang.org/pkg/",
            "https://golang.org/cmd/",
        },
    },
    "https://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "https://golang.org/",
            "https://golang.org/cmd/",
            "https://golang.org/pkg/fmt/",
            "https://golang.org/pkg/os/",
        },
    },
    "https://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "https://golang.org/",
            "https://golang.org/pkg/",
        },
    },
    "https://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "https://golang.org/",
            "https://golang.org/pkg/",
        },
    },
}

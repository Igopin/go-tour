package main

import ( 
    "fmt"
    "strings"
    "strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (addr IPAddr) String() string {
    //return fmt.Sprintf("%v.%v.%v.%v", addr[0], addr[1], addr[2], addr[3])
    
    res := make([]string, len(addr))
    for i, v := range addr {
        res[i] = strconv.Itoa(int(v))
    }
    
    return strings.Join(res, ".")
}

func main() {
    hosts := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}
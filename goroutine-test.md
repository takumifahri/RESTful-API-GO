```go
package main

import (
    "fmt"
    "time"
    "math/rand"
    "sync"
)

func testPrints(id int) {
    for i := 0; i < 10; i++ {
        fmt.Printf("[%d] Hello World %d\n", id, i)
        amt := time.Duration(rand.Intn(1000))
        time.Sleep(amt * time.Millisecond)
    }
}

func main() {
    var sharedRecourse string
    var wg sync.WaitGroup
    var mutex sync.Mutex
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            mutex.Lock()
            fmt.Println("previous value of sharedRecourse: ", sharedRecourse)
            sharedRecourse = fmt.Sprintf("Hello World [%d]", i)
            fmt.Println("new value of sharedRecourse:", sharedRecourse)
            mutex.Unlock()
            // testPrints(i)
        }(i)
    }
    wg.Wait()

    fmt.Printf("Final value of sharedRecourse: %s\n", sharedRecourse)
}
```

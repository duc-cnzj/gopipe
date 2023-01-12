# go pipeline

## usage

```go
package main

import (
  "fmt"
  "github.com/duc-cnzj/gopipe"
)

func main()  {
    gopipe.NewPipeline[any]().
        Send(nil).
        Through(func(t any, next func()) {
            fmt.Println(1)
            next()
            fmt.Println(2)
        }).
        Through(func(t any, next func()) {
            fmt.Println(3)
            next()
            fmt.Println(4)
        }).
        Then(func(any) {
            fmt.Println("core logic")
        })
    // Output:
    // 1
    // 3
    // core logic
    // 4
    // 2
}
```

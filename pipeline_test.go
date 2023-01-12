package gopipe

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"
)

func ExampleNewPipeline() {
    NewPipeline[any]().
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

func TestMyPipeline_Send(t *testing.T) {
    var result []string
    NewPipeline[string]().
        Send("xxx").
        Through(func(msg string, next func()) {
            assert.Equal(t, "xxx", msg)
            result = append(result, "1")
            result = append(result, "2")
            next()
            result = append(result, "3")
        }, func(msg string, next func()) {
            assert.Equal(t, "xxx", msg)
            result = append(result, "4")
            result = append(result, "5")
            next()
            result = append(result, "6")
        }).
        Then(func(msg string) {
            result = append(result, msg)
        })

    assert.Equal(t, []string{"1", "2", "4", "5", "xxx", "6", "3"}, result)

    called := false
    NewPipeline[int]().Then(func(int) {
        called = true
    })
    assert.True(t, called)

    type obj struct {
        name string
    }

    oo := &obj{name: "app"}
    NewPipeline[*obj]().
        Send(oo).
        Through(func(o *obj, next func()) {
            o.name += "1"
            next()
            o.name += "2"
        }).
        Through(func(o *obj, next func()) {
            o.name += "3"
            next()
            o.name += "4"
        }).
        Then(func(o *obj) {
            o.name += "base"
        })

    assert.Equal(t, "app13base42", oo.name)
}

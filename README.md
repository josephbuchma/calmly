# Calmly
[![Build Status](https://travis-ci.org/JosephBuchma/calmly.svg?branch=master)](https://travis-ci.org/JosephBuchma/calmly)

Package calmly is example of golang package that should not exist.

```go
import (
  "fmt"
  "github.com/JosephBuchma/calmly"
)

type MyError struct{}
type MyAnotherError struct{}

func main(){
  calmly.Try(func(){
    panic(MyError{}) // `raise` error (any type)

  // Ideally you should pass into Catch only type, (e.g. `MyError`, not `MyError{}`),
  // but such syntax is not supported by golang yet.
  // Discussion is here https://groups.google.com/forum/#!topic/golang-nuts/dYMlhyq5FpA
  }).Catch(MyError{}, func(e calmly.E){
    fmt.Println("Error catched")

  // Add as many catches as you need...  
  }).Catch(MyAnotherError{}, func(e calmly.E) {
    fmt.Println("Not this time")

  // e.g. `catch (...)`
  }).CatchAny(func(e calmly.E) {
    fmt.Println("Not this time")

  // Finally runs anyway. Finally is required clause.
  // If you don't need any finalization, simply pass nil as parameter instead of func(). 
  }).Finally(func(){
    fmt.Println("Finalization")
  })
}
```

See calmly_test.go for more examples.

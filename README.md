# trie library by golang
### Usage
    go get -u github.com/kovey/trie-go
### Examples
```golang
    package main

    import (
        "github.com/kovey/trie-go/filter"
    )

    func main() {
        if err := filter.Load("/path/to/keywords", ","); err != nil {
            panic(err)
        }

        if filter.Has("test") {
            fmt.Print("yes")
        } else {
            fmt.Print("no")
        }

        fmt.Print(filter.Replace("test"))
    }

```

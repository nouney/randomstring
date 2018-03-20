# Randomstring

A small Golang package to easely generate random strings.

The code is based on this awesome [StackOverflow answer](https://stackoverflow.com/a/31832326/2432477) by [icza](https://stackoverflow.com/users/1705598/icza).

## Installation

To install `randomstring`, run the following command:

 `$ go get github.com/nouney/randomstring`

## Example

```golang
package main

import (
    "fmt"

    "github.com/nouney/randomstring"
)

func main() {
    // Generate an alphanum random string of 16 chars
    str := randomstring.Generate(16)
    fmt.Println("Alphanum:", str)

    rsg, err := randomstring.NewGenerator(randomstring.CharsetNum)
    if err != nil {
        log.Panic(err)
    }
    str = rsg.Generate(6)
    fmt.Println("Num:", str)

    rsg, err = randomstring.NewGenerator("AbCdEfGhIjKlMnOpQrStUvWxYz")
    if err != nil {
        log.Panic(err)
    }
    str = rsg.Generate(6)
    fmt.Println("Custom charset:", str)
}

```

## Default charsets

- `CharsetAlphaNum`: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
- `CharsetNum`: "0123456789"
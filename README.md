# moerr

Minimal error creation utility for Go.  
Provides simple functions for constructing static and formatted errors with a clean and consistent API.

## Installation

```sh
go get github.com/MaiWittawat/moerr-golang
```


## Why moerr?

* Centralizes error creation for projects that want a dedicated error factory

* Keeps code consistent by avoiding mixed usage of errors.New and fmt.Errorf

* Simple, zero-dependency, idiomatic Go

## Using

``` go
import "github.com/MaiWittawat/moerr-golang"

func main() {
    err1 := moerr.New("something went wrong")
    err2 := moerr.NewErrorf("invalid value: %v", 42)

    fmt.Println(err1)
    fmt.Println(err2)
}
```

## API Overview

```func New(text string) error```

Creates a simple error with a static message.

``` go
err := moerr.New("user not found")
```

```func NewErrorf(format string, a ...any) error``` <br>
Creates a formatted error message.
Equivalent to fmt.Errorf(format, a...) but using moerr's internal error type.
``` go
err := moerr.NewErrorf("invalid id: %d", id)
```

## Project Structure

``` sh
.
├── go.mod
├── main.go        (usage example)
└── moerr
    └── moerr.go   (library source)
```

## License
MIT — feel free to use it in personal or commercial projects.
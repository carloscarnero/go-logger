> This distribution provides a *very* simple and slightly opinionated Go
> package helpful to create [slog.Logger][LOGGER] instances. I found myself
> using the same patterns repeatedly, and decided to lift the code into
> this independent package. Unless you, somehow, find this convenient,
> there is probably no reason why you should actually use this 🤷

![License](https://img.shields.io/badge/License-MIT-blue.svg)
[![codecov](https://codecov.io/github/carloscarnero/go-logger/graph/badge.svg?token=Bg1xghG6HS)](https://codecov.io/github/carloscarnero/go-logger)

Retrieve the package with:

```console
go get -u -v go.carloscarnero.stream/go-logger
```

and then it will be ready to use in your project:

```go
package main

import (
    "os"

    "go.carloscarnero.stream/go-logger"
)

func main() {
    // l will be a pointer to a standard library slog.Logger instance. Note
    // that no error handling is included in this example.
    l, _ := logger.New(os.Stdout, "TEXT", "INFO", false)

    l.Info("Hello, world!")
}
```

Running the above program will send `Hello, world!` to the standard output:

```console
level=INFO msg="Hello, world!"
```

> [!NOTE]
> While the module path is `go.carloscarnero.stream/go-logger`, the package
> is called `logger`. This is not very elegant but results from having to
> name the source code repository in that particular way.

[LOGGER]: https://pkg.go.dev/log/slog#Logger

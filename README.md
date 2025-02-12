> This distribution provides a *very* simple and slightly opinionated Go
> package helpful to create [slog.Logger][LOGGER] instances. I found myself
> using the same patterns repeatedly, and decided to lift the code into
> this independent package. Unless you, somehow, find this convenient,
> there is probably no reason why you should actually use this 🤷

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/license/mit)
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

The first parameter to the `New` method is the writer where all output will
be sent to, which is nothing more than a `io.Writer`. This can also be used
to suppress all output if the parameter is either `nil` or `io.Discard`.

The second parameter is the logging format, which can be either `JSON` or
`TEXT`. The third parameter is the logging level, which can be either
`DEBUG`, `INFO`, `WARN`, `ERROR`, or `NONE`. This last one will also
suppress all output.

> Why two mechanisms to supress or discard output? In many cases, defining
> the actual writer is a source code concern and maybe not very flexible.
> On the other hand, defining the level may be easier as it is common to
> have this value defined in environment variables, for instance.

The fourth parameter allows to supress timestamps in the output, which can
be useful in debugging, development, and testing scenarios.

> [!NOTE]
> While the module path is `go.carloscarnero.stream/go-logger`, the package
> is called `logger`. This is not very elegant but results from having to
> name the source code repository in that particular way.

[LOGGER]: https://pkg.go.dev/log/slog#Logger

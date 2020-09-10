# gosh
Directly execute Go source files with shebang lines

## Installation

```
go get github.com/mihs/gosh
```

## Usage

Add a shebang line to a go source file, like in the following "example.go" file.

```go
#!/usr/bin/env gosh
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
```

Then make the file executable and run it.


```
$ chmod +x example.go
$ ./example
Hello world!
```

If the source file starts with `#!`, gosh will create a temporary file in your temporary directory (/tmp)
with the first line commented out so that the go compiler will not complain.

## Alternatives

For more features, see [gorun](https://github.com/erning/gorun).

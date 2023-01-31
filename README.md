# stripmargin

A tiny Go library that allows you to write multiline strings and strip margin, just like you would in Scala.

## Usage

```go
import (
  "fmt"
  "github.com/cstanze/stripmargin"
)

// Use '|' to set the left margin,
// then surround in stripmargin.StripMargin
fmt.Println(stripmargin.StripMargin(`
  |Hello,
  |  world!`))

// Prints:
//
// Hello,
//   world!

```

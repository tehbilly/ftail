# File tail

A very simple wrapper around [https://github.com/ActiveState/tail](github.com/ActiveState/tail)
that I use mainly as a convenience lib.

## Example

```go
package main

import (
	"fmt"
	"os"

	"github.com/tehbilly/ftail"
)

func main() {
  // Should be safe to call regardless. Needs to be called 'just in case'
	defer ftail.Cleanup()
	lines, err := ftail.TailFile("somefile.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for line := range lines {
		fmt.Printf("%+v\n", line)
	}
}
```

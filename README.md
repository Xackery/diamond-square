# Go Diamond-Square Algorithmn
[Diamond-Square Algorithmn](https://en.wikipedia.org/wiki/Diamond-square_algorithm) is a method for making heighmaps and terrain generation stuff. I doubt my implementation is perfect, but it is inspired by a couple examples I've seen online. If you improve it, feel free to fork and merge. 

## Usage
````bash
go get github.com/xackery/diamondsquare
````

````
package main
import (
	"github.com/xackery/diamondsquare"
	"fmt"
)
func main() {
	data, e := diamondsquare.Generate(257, 3883, 500)
	if e != nil {
		fmt.Println(e.Error())
	}
}
}
````
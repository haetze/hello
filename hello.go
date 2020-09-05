package main // import "github.com/you/hello"

import (
	"fmt"
	"rsc.io/quote"
	"github.com/haetze/hello/me"
)

func main() {
	fmt.Println(quote.Hello())
	fmt.Println(me.Name())
}

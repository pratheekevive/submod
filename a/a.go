package main

import (
	"fmt"

	"github.com/pratheekevive/submod/b"
	"github.com/pratheekevive/submod/c"
)

const Name = b.Name

func main() {
	fmt.Println(Name)

	c.Hello()
}

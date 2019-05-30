package main

import (
	"fmt"
	"github.com/karalabe/hid"
)

func main() {

	info := hid.Enumerate(0,0)
	fmt.Print(info)
}

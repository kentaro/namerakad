package main

import (
	"github.com/kentaro/namerakad/client"
)

func main() {
	c := namerakad.NewClient()
	c.Run()
}

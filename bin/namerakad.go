package main

import (
	"github.com/kentaro/namerakad/server"
)

func main() {
	s := namerakad.NewServer()
	s.Run()
}

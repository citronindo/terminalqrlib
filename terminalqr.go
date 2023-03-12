package main

import "C"

import (
	"os"

	"github.com/mdp/qrterminal/v3"
)

func main() {}

// printqr :
//
//export printqr
func printqr(code *C.char) {
	s := C.GoString(code)
	qrterminal.Generate(s, qrterminal.L, os.Stdout)
}

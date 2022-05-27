package main

import "C"

import (
	"os"

	"github.com/mdp/qrterminal"
)

func main() {}

//printqr :
//export printqr
func printqr(code *C.char) {
	s := C.GoString(code)
	qrterminal.GenerateHalfBlock(s, qrterminal.L, os.Stdout)
}

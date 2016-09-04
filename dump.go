package main

import (
	"fmt"
	"io"
	"os"
)

func doDumpList(list string) {
	r := internalListReader(list)
	if r == nil {
		fmt.Fprintf(os.Stderr, "error: list %q does not exist\n", list)
		return
	}
	defer r.Close()
	io.Copy(os.Stdout, r)
}

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	out := os.Stdout
	out.WriteString(`[]byte{`)

	buf := make([]byte, 64*1024*1024)
	for {
		n, err := f.Read(buf)
		for _, b := range buf[:n] {
			fmt.Fprintf(out, "0x%02x, ", b)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}

	out.WriteString(`}`)
}

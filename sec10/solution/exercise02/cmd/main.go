package main

import (
	"fmt"
	"io"
	"os"

	"github.com/egbordzor/learn-glft/sec10/solution/exercise02/ms"
	"github.com/egbordzor/learn-glft/sec10/solution/exercise02/rw"
)

func main() {
	var m, _ = ms.NewMemStore(200)
	writeRecords(m)
	io.Copy(os.Stdout, m)
	fmt.Println()
}

func writeRecords(w io.Writer) {
	w, _ = rw.NewRecordWriter(w)
	io.WriteString(w, "Hello, world")
	io.WriteString(w, ". ")
	io.WriteString(w, "It is a wonderful day.")
	fmt.Println(w)
}

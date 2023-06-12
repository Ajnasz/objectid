package main

import (
	"flag"
	"fmt"

	"github.com/Ajnasz/objectid"
)

func main() {
	n := flag.Int("n", 1, "number of objectid to generate")
	separator := flag.String("s", "\n", "separator between objectids")
	flag.Parse()

	for i := 0; i < *n; i++ {
		oid := objectid.New()
		fmt.Printf("%s%s", oid, *separator)
	}
}

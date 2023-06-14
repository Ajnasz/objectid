package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Ajnasz/objectid"
)

func printTimeOfObjectID(idString string) {
	oid, err := objectid.FromHex(idString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(oid.Time())
}

func generate(n int, separator string) {
	for i := 0; i < n; i++ {
		oid := objectid.New()
		fmt.Printf("%s%s", oid, separator)
	}
}

func main() {
	n := flag.Int("n", 1, "number of objectid to generate")
	separator := flag.String("s", "\n", "separator between objectids")
	toTime := flag.String("t", "", "convert objectid to time")
	flag.Parse()

	if toTime != nil && *toTime != "" {
		printTimeOfObjectID(*toTime)
		return
	}

	generate(*n, *separator)
}

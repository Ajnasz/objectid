package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Ajnasz/objectid"
)

func main() {
	n := flag.Int("n", 1, "number of objectid to generate")
	separator := flag.String("s", "\n", "separator between objectids")
	toTime := flag.String("t", "", "convert objectid to time")
	flag.Parse()

	if toTime != nil && *toTime != "" {
		oid, err := objectid.FromHex(*toTime)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(oid.Time())
		return
	}

	for i := 0; i < *n; i++ {
		oid := objectid.New()
		fmt.Printf("%s%s", oid, *separator)
	}
}

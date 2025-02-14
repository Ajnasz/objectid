package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Ajnasz/objectid"
)

var (
	version string
	build   string
)

func printTimeOfObjectID(idString string) {
	oid, err := objectid.FromHex(idString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(oid.Time())
}

func printHex(oid objectid.ObjectID, separator string) {
	fmt.Printf("%s%s", oid.Hex(), separator)
}

func printBase64(oid objectid.ObjectID, separator string) {
	fmt.Printf("%s%s", oid.Base64(), separator)
}

func generate(n int, separator string, format string) {
	var printer func(objectid.ObjectID, string)
	switch format {
	case "hex":
		printer = printHex
	case "base64":
		printer = printBase64
	default:
		printer = printHex
	}
	for i := 0; i < n; i++ {
		oid := objectid.New()
		printer(oid, separator)
	}
}

func generateFromDate(date string, format string) {
	var printer func(objectid.ObjectID, string)
	switch format {
	case "hex":
		printer = printHex
	case "base64":
		printer = printBase64
	default:
		printer = printHex
	}

	oid, err := objectid.FromTime(date)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	printer(oid, "")
}

func main() {
	n := flag.Int("n", 1, "number of objectid to generate")
	getVersion := flag.Bool("version", false, "print version")
	separator := flag.String("separator", "\n", "separator between objectids")
	format := flag.String("format", "hex", "format of objectid: hex, base64")
	toTime := flag.String("to-time", "", "convert objectid to time. Use - to read from stdin")
	fromTime := flag.String("from-time", "", "create a new objectid from a date time (RFC3339, $(date -I), $(date -Ihours), $(date -d -Iminutes), $(date -Iseconds). Use - to read from stdin")
	flag.Parse()

	if *getVersion {
		fmt.Printf("%s %s", version, build)
		fmt.Println()
		os.Exit(0)
	}

	if toTime != nil && *toTime != "" {
		timeToConvert := *toTime

		if *toTime == "-" {
			fmt.Scanln(&timeToConvert)
		}

		printTimeOfObjectID(timeToConvert)
		return
	}

	if *fromTime != "" {
		date := *fromTime

		if *fromTime == "-" {
			fmt.Scanln(&date)
		}

		generateFromDate(date, *format)
		return
	}

	generate(*n, *separator, *format)
}

package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

const timeFormat = "2006-01-02 15:04:05.999999999 -0700 MST"

func main() {
	testMap := map[string]string{
		"Australia/Brisbane": "2020-10-29 15:30:00 +1000 AEST",
		"Australia/Sydney":   "2020-10-29 15:30:00 +1100 AEDT",
		"Europe/Berlin":      "2020-10-29 15:30:00 +0100 CET",
	}

	for location, expectedDate := range testMap {
		loc, err := time.LoadLocation(location)
		if err != nil {
			log.Printf("unable to load location %s: %v", location, err)
			continue
		}
		p, err := time.Parse(timeFormat, expectedDate)
		if err != nil {
			log.Printf("unable to parse date/time %s: %v", location, err)
			continue
		}

		d := time.Date(p.Year(), p.Month(), p.Day(), p.Hour(), p.Minute(), p.Second(), p.Nanosecond(), loc)

		if d.String() == expectedDate {
			log.Printf("Expected time stamp returned for %s", location)
		} else {
			log.Printf("Incorrect time stamp returned for %s", location)
		}
		log.Printf("\tExpected: %s", expectedDate)
		log.Printf("\tActual:   %s", d.String())
	}
}

func failErr(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}

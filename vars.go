package main

import "sync"

const (
	// JSONregex for printing json
	JSONregex = `(\{.*\})+`

	// outputRegex for printing yaml/json output
	outputJSONregex = `(?m)^[ ]*([^\r\n:]+?)\s*:`

	// statment regex
	stateRegex = `^START|END|REPORT`
)

var (
	profile string
	region  string
	debug   bool
	colors  bool
	wg      sync.WaitGroup

	// define logger
	log = logger{
		DebugMode: &debug,
		Colors:    &colors,
	}
)

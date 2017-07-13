package main

import "sync"

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

// Copyright 2012 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"launchpad.net/goamz/ec2"
	"launchpad.net/goamz/ec2/ec2test"
	"log"
)

var initialState = flag.String("initial", "", "Initial instance state.")

func init() {
	flag.Parse()
}

func main() {
	server, err := ec2test.NewServer()
	if err != nil {
		log.Fatalf("Failed to start server: %s.", err)
	}
	defer server.Quit()
	if *initialState != "" {
		var code int
		switch *initialState {
		case "pending":
			code = 0
		case "running":
			code = 16
		case "shutting-down":
			code = 32
		case "terminated":
			code = 48
		case "stopping":
			code = 64
		case "stopped":
			code = 80
		default:
			log.Fatalf("Invalid initial state: %q.\nSee http://goo.gl/y3ZBq for a list of valid states.", *initialState)
		}
		state := ec2.InstanceState{Code: code, Name: *initialState}
		server.SetInitialInstanceState(state)
	}
	fmt.Printf("Server is listening at %s.\nPress Ctrl-C to close it.\n", server.URL())
	for {
	}
}

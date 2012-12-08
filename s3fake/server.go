// Copyright 2012 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"launchpad.net/goamz/s3/s3test"
	"log"
)

func main() {
	server, err := s3test.NewServer()
	if err != nil {
		log.Fatalf("Failed to start server: %s.", err)
	}
	defer server.Quit()
	fmt.Printf("Server is listening at %s.\nPress Ctrl-C to close it.\n", server.URL())
	for {
	}
}

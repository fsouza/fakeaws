// Copyright 2012 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.package main

import (
	"fmt"
	"github.com/fsouza/go-iam/iamtest"
	"log"
)

func main() {
	server, err := iamtest.NewServer()
	if err != nil {
		log.Fatalf("Failed to start server: %s.", err)
	}
	defer server.Quit()
	fmt.Printf("Server is listening at %s.\nPress Ctrl-C to close it.\n", server.URL())
	for {
	}
}

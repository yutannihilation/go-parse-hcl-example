package main

/**

c.f.
- godoc:
  - https://godoc.org/github.com/hashicorp/hcl2/hclparse
  - https://godoc.org/github.com/hashicorp/hcl2/gohcl
- parsing: https://github.com/hashicorp/hcl2/blob/master/guide/go_parsing.rst
- decoding: https://github.com/hashicorp/hcl2/blob/master/guide/go_decoding_gohcl.rst
**/

import (
	"fmt"
	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
	"log"
)

type foo struct {
	A string   `hcl:"a"`
	B []string `hcl:"b"`
}

func main() {
	parser := hclparse.NewParser()
	f, parseDiags := parser.ParseHCLFile("test.hcl")
	if parseDiags.HasErrors() {
		log.Fatal(parseDiags.Error())
	}

	var fooInstance foo
	decodeDiags := gohcl.DecodeBody(f.Body, nil, &fooInstance)
	if decodeDiags.HasErrors() {
		log.Fatal(decodeDiags.Error())
	}

	fmt.Printf("%#v", fooInstance)
}

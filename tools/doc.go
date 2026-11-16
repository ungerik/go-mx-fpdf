// Package tools pins the build-time tooling used to maintain go-mx-fpdf — the
// ungerik/go-enum code generator for the methods of the //#enum types — in a
// separate nested module, so go-enum and its dependency tree are not part of
// the github.com/ungerik/go-mx-fpdf module that consumers import.
//
// Generation is driven by //go:generate directives in the files that define
// enums:
//
//	//go:generate go -C tools tool go-enum ../$GOFILE
//
// CI checks that the committed methods are up to date with
//
//	go -C tools tool go-enum -validate ../...
package tools

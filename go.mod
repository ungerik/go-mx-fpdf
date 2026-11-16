module github.com/ungerik/go-mx-fpdf

go 1.27.0

require (
	codeberg.org/go-pdf/fpdf v0.12.0
	github.com/domonda/go-pretty v1.0.0
	// The parity tests compare this legacy wrapper's output against the native
	// github.com/ungerik/go-mx/pdf package.
	github.com/ungerik/go-mx v0.0.0-20261115173550-0e6784d998ed
)

require github.com/domonda/go-errs v1.1.0 // indirect

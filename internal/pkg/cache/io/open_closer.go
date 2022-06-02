package io

import "io"

// OpenCloser -
type OpenCloser interface {
	Opener
	io.Closer
}

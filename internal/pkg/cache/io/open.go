package io

// Opener -
type Opener interface {
	Open(dsn string) error
}

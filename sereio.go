// Package sereio contents everything
package sereio

var SOA map[string]interface{} // collection of services, storages, etc

type DataFrame interface {
	BehaviorSet()
	Views()
	Select()    // select view by name
	Converter() // term-id extern (local?) converter
	Stat()      // select sere/stat of Select()ed view
}

// Request extern source
func NewGet(url string) {

}

// Read local file
func NewRead(path string) {

}

type Error struct {
	pkg string // package with param
	err string // error.New() predefined in package
}

// Collect package error and make []string
func ErrorList(err []*Error) {
	// Collect errors
	// Make log
}

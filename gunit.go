// Package gunit provides "testing" package hooks and convenience
// functions for writing tests in an xUnit style.
// NOTE: Only some of the exported names in this package
// are meant to be referenced by users of this package:
//
// - Fixture              // (as an embedded field on your xUnit-style struct)
// - Fixture.So(...)      // (as a convenient assertion method: So(expected, should.Equal, actual))
// - Fixture.Ok(...)      // (as a convenient boolean assertion method: Ok(condition, optionalMessage))
// - Fixture.Error(...)   // (works just like *testing.T.Error(...))
// - Fixture.Errorf(...)  // (works just like *testing.T.Errorf(...))
// - Fixture.Print(...)   // (works just like fmt.Print)
// - Fixture.Printf(...)  // (works just like fmt.Printf)
// - Fixture.Println(...) // (works just like fmt.Println)
//
// The rest are called from code generated by the command at
// github.com/smartystreets/gunit/gunit.
// Please see the README file and the examples folder for examples.
package gunit

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"

	"github.com/smartystreets/assertions"
)

// TT represents the functional subset from *testing.T needed by Fixture.
type TT interface {
	Log(args ...interface{})
	Fail()
	Failed() bool
	SkipNow()
	Parallel()
}

// Fixture keeps track of test status (failed, passed, skipped) and
// handles custom logging for xUnit style tests as an embedded field.
type Fixture struct {
	t       TT
	log     *bytes.Buffer
	verbose bool
}

// NewFixture is called by generated code.
func NewFixture(t TT, verbose bool) *Fixture {
	return &Fixture{t: t, verbose: verbose, log: &bytes.Buffer{}}
}

// Parallel is analogous to *testing.T.Parallel.
func (self *Fixture) Parallel() {
	self.t.Parallel()
}

// So is a convenience method for reporting assertion failure messages,
// say from the assertion functions found in github.com/smartystreets/assertions/should.
// Example: self.So(actual, should.Equal, expected)
func (self *Fixture) So(actual interface{}, assert func(actual interface{}, expected ...interface{}) string, expected ...interface{}) bool {
	ok, failure := assertions.So(actual, assert, expected...)
	if !ok {
		self.t.Fail()
		self.reportFailure(failure)
	}
	return ok
}

func (self *Fixture) Ok(condition bool, messages ...string) {
	if !condition {
		if len(messages) == 0 {
			messages = append(messages, "Expected condition to be true, was false instead.")
		}
		self.t.Fail()
		self.reportFailure(strings.Join(messages, ", "))
	}
}

func (self *Fixture) Error(args ...interface{}) {
	self.t.Fail()
	self.reportFailure(fmt.Sprint(args...))
}

func (self *Fixture) Errorf(format string, args ...interface{}) {
	self.t.Fail()
	self.reportFailure(fmt.Sprintf(format, args...))
}

func (self *Fixture) reportFailure(failure string) {
	self.Print(NewFailureReport(failure))
}

// Print is analogous to fmt.Print and is ideal for printing in the middle of a test case.
func (self *Fixture) Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(self, a...)
}

// Printf is analogous to fmt.Printf and is ideal for printing in the middle of a test case.
func (self *Fixture) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(self, format, a...)
}

// Println is analogous to fmt.Println and is ideal for printing in the middle of a test case.
func (self *Fixture) Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(self, a...)
}

// Write implements io.Writer.
func (self *Fixture) Write(p []byte) (int, error) {
	return self.log.Write(p)
}

// Finalize is called by generated code.
func (self *Fixture) Finalize() {
	if r := recover(); r != nil {
		self.recover(r)
	}

	if self.t.Failed() || (self.verbose && self.log.Len() > 0) {
		self.t.Log("\n" + strings.TrimSpace(self.log.String()) + "\n")
	}
}

// Failed is analogous to *testing.T.Failed().
func (self *Fixture) Failed() bool {
	return self.t.Failed()
}

func (self *Fixture) recover(r interface{}) {
	self.Println("X PANIC:", r)
	buffer := make([]byte, 1024*16)
	runtime.Stack(buffer, false)
	self.Println(strings.TrimSpace(string(buffer)))
	self.Println("* (Additional tests may have been skipped as a result of the panic shown above.)")
	self.t.Fail()
}

package gunit

// testingT represents the functional subset from *testing.T needed by Fixture.
type testingT interface {
	Name() string
	Log(args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatalf(format string, args ...interface{})
}

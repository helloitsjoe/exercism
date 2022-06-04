package erratum

// import "io"

// These are the support types and interface definitions used in the
// implementation if your Use function. See the test suite file at
// for information on the expected implementation.
//
// Because this is part of the package "erratum", if your solution file
// is also declared in the package you will automatically have access to
// these definitions (you do not have to re-declare them).

// TransientError is an error that may occur while opening a resource via
// ResourceOpener.
type TransientError struct {
	err error
}

func (e TransientError) Error() string {
	return e.err.Error()
}

type FrobError struct {
	defrobTag string
	err       error
}

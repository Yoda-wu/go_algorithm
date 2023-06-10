package util_test

import "time"

var (
	// true: test only one case in AssertEqualRunResultsInf / CheckRunResultsInf
	Once bool

	// true: only print test case number when assertion failed
	DisableLogInput bool

	// when DebugTLE > 0, a running case would cause a fatal error when timeout
	DebugTLE = 2 * time.Second
)

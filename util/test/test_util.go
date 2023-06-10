package util_test

import (
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type ioFunc func(io.Reader, io.Writer)

func isTLE(f func()) bool {
	if DebugTLE == 0 {
		f()
		return false
	}

	done := make(chan struct{})
	timer := time.NewTimer(DebugTLE)
	defer timer.Stop()
	go func() {
		defer close(done)
		f()
	}()
	select {
	case <-done:
		return false
	case <-timer.C:
		return true
	}
}

func AssertEqualString(t *testing.T, testCases [][2]string, runFunc ioFunc) {
	if len(testCases) == 0 {
		return
	}

	allPass := true
	for curCaseNum, tc := range testCases {
		input := removeExtraSpace(tc[0])

		expectOuput := removeExtraSpace(tc[1])

		mockReader := strings.NewReader(input)
		mockWriter := &strings.Builder{}
		f := func() {
			runFunc(mockReader, mockWriter)
		}
		if isTLE(f) {
			allPass = false
			t.Errorf("Time Limit Exceeded %d\nInput:\n%s", curCaseNum+1, input)
			continue
		} else {
			f()
		}

		actualOutput := removeExtraSpace(mockWriter.String())
		t.Run(fmt.Sprintf("Case %d", curCaseNum+1), func(t *testing.T) {
			if !assert.Equal(t, expectOuput, actualOutput, "Wrong Answer %d\nInput:\n%s", curCaseNum+1, input) {
				allPass = false
				handleOutput(t, actualOutput)
			}
		})

	}
	if !allPass {
		return
	}
	t.Log("ok!")
}

func AssertEqualCase(t *testing.T, rawText string, runFunc ioFunc){
	
}


// -------------------- helper --------------------
func removeExtraSpace(s string) string {
	s = strings.TrimSpace(s)
	sp := strings.Split(s, "\n")
	for i := range sp {
		sp[i] = strings.TrimSpace(sp[i])
	}
	return strings.Join(sp, "\n")
}

func handleOutput(t *testing.T, s string) {
	t.Logf("actual output = %s\n", s)
}

package cf_daily

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCF480C(t *testing.T) {

	input := "5 2 4 1"
	mockReader := strings.NewReader(input)
	mockWriter := &strings.Builder{}
	CF480C(mockReader, mockWriter)
	expectOutput := "2"
	actualOutput := mockWriter.String()
	assert.Equal(t, expectOutput, actualOutput)
}

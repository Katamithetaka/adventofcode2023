package day1_part2;

import (
	_ "embed"
	"testing"
    "gotest.tools/assert"
)

//go:embed test.txt
var test string

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func Test(t *testing.T) {
    t.Helper()

    expected_value := "281"
    result := Run(test)

    assert.Equal(t, result, expected_value)

}
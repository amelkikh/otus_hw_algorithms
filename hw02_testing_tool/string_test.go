package hw02_testing_tool

import (
	"os"
	"path"
	"strconv"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

type StringLength struct {
}

func (s *StringLength) Run(str []string) string {
	return strconv.Itoa(utf8.RuneCount([]byte(str[0])))
}

func TestNewTesterStringLength(t *testing.T) {
	dir, err := os.Getwd()
	require.NoError(t, err)
	tester := NewTester(&StringLength{}, path.Join(dir, "tests", "0.String"))
	err = tester.RunTests()
	require.NoError(t, err)
}

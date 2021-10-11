package hw02_testing_tool

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type LuckyTickets struct {
}

func (s *LuckyTickets) Run(str []string) string {
	n, err := strconv.Atoi(str[0])
	if err != nil {
		log.Fatal(fmt.Errorf("invalid input integer value: %v", str[0]))
	}
	if n <= 0 {
		log.Fatal(fmt.Errorf("value should be positive: %v", str[0]))
	}

	basis := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	var rows int
	for iter := 0; iter < n; iter++ {
		if iter < 1 {
			rows = 1
		} else {
			rows = 10
		}
		cols := (iter+1)*9 + 1

		m := make(map[int][]int, rows)
		for i := 0; i < rows; i++ {
			row := make([]int, 0, cols)
			row = append(row, make([]int, i)...)
			row = append(row, basis...)
			row = append(row, make([]int, cols-len(row))...)
			m[i] = row
		}
		r := make([]int, cols)
		for k := 0; k < len(m); k++ {
			for j := 0; j < cols; j++ {
				r[j] += m[k][j]
			}
		}

		basis = r
	}

	sum := 0
	for _, v := range basis {
		sum += v * v
	}

	return strconv.Itoa(sum)
}

func TestNewTesterLuckyTickets(t *testing.T) {
	dir, err := os.Getwd()
	require.NoError(t, err)
	tester := NewTester(&LuckyTickets{}, path.Join(dir, "tests", "1.Tickets"))
	err = tester.RunTests()
	require.NoError(t, err)
}

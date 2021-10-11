package hw03_bit_arithmetic

import (
	"fmt"
	"log"
	"math/bits"
	"os"
	"path"
	"strconv"
	"testing"

	tt "github.com/amelkikh/otus_hw_algorithms/hw02_testing_tool"
	"github.com/stretchr/testify/require"
)

type BitKing struct{}

const (
	KingNoA = 0xfefefefefefefefe
	KingNoH = 0x7f7f7f7f7f7f7f7f
)

func (q *BitKing) Run(strings []string) string {
	p, err := strconv.Atoi(strings[0])
	if err != nil {
		log.Fatal(fmt.Errorf("parse position: %w", err))
	}
	var K uint64
	K = 1 << p
	Ka := K & KingNoA
	Kh := K & KingNoH
	M := (Ka << 7) | (K << 8) | (Kh << 9) |
		(Ka >> 1) | (Kh << 1) |
		(Ka >> 9) | (K >> 8) | (Kh >> 7)

	pop := bits.OnesCount64(M)
	return fmt.Sprintf("%d\n%d", pop, M)
}

func TestBitKing(t *testing.T) {
	dir, err := os.Getwd()
	require.NoError(t, err)

	tester := tt.NewTester(&BitKing{}, path.Join(dir, "tests", "1.Bitboard - Король"))
	err = tester.RunTests()
	require.NoError(t, err)
}

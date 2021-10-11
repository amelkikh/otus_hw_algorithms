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

type BitQueen struct{}

const (
	QueenNoA = 0xfefefefefefefefe
	QueenNoH = 0x7f7f7f7f7f7f7f7f
)

func (q *BitQueen) Run(strings []string) string {
	p, err := strconv.Atoi(strings[0])
	if err != nil {
		log.Fatal(fmt.Errorf("parse position: %w", err))
	}

	var B, M uint64
	B = 1 << p
	// SouthWest
	for K := B; K > 0; {
		K = (K & BishopNoA) >> 9
		M += K
	}
	// NorthWest
	for K := B; K > 0; {
		K = (K & BishopNoA) << 7
		M += K
	}
	// SouthEast
	for K := B; K > 0; {
		K = (K & BishopNoH) >> 7
		M += K
	}
	// NorthEast
	for K := B; K > 0; {
		K = (K & BishopNoH) << 9
		M += K
	}
	// East
	for K := B; K > 0; {
		K = (K & BishopNoH) << 1
		M += K
	}
	// West
	for K := B; K > 0; {
		K = (K & BishopNoA) >> 1
		M += K
	}
	// North
	for K := B; K > 0; {
		K = K << 8
		M += K
	}
	// South
	for K := B; K > 0; {
		K = K >> 8
		M += K
	}

	pop := bits.OnesCount64(M)
	return fmt.Sprintf("%d\n%d", pop, M)
}

func TestBitQueen(t *testing.T) {
	dir, err := os.Getwd()
	require.NoError(t, err)
	tester := tt.NewTester(&BitQueen{}, path.Join(dir, "tests", "5.Bitboard - Ферзь"))
	err = tester.RunTests()
	require.NoError(t, err)
}

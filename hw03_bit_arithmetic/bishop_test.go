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

type BitBishop struct{}

const (
	BishopNoA = 0xfefefefefefefefe
	BishopNoH = 0x7f7f7f7f7f7f7f7f
)

func (e *BitBishop) Run(strings []string) string {
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

	pop := bits.OnesCount64(M)
	return fmt.Sprintf("%d\n%d", pop, M)
}

func TestBitBishop(t *testing.T) {
	dir, err := os.Getwd()
	require.NoError(t, err)
	tester := tt.NewTester(&BitBishop{}, path.Join(dir, "tests", "4.Bitboard - Слон"))
	err = tester.RunTests()
	require.NoError(t, err)
}

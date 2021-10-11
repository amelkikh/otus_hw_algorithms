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

type BitRook struct{}

const (
	RookNoA = 0xfefefefefefefefe
	RookNoH = 0x7f7f7f7f7f7f7f7f
)

func (r *BitRook) Run(strings []string) string {
	p, err := strconv.Atoi(strings[0])
	if err != nil {
		log.Fatal(fmt.Errorf("parse position: %w", err))
	}

	var B, M uint64
	B = 1 << p
	// East
	for K := B; K > 0; {
		K = (K & RookNoH) << 1
		M += K
	}
	// West
	for K := B; K > 0; {
		K = (K & RookNoA) >> 1
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

func TestBitRook(t *testing.T) {
	dir, err := os.Getwd()
	require.NoError(t, err)
	tester := tt.NewTester(&BitRook{}, path.Join(dir, "tests", "3.Bitboard - Ладья"))
	err = tester.RunTests()
	require.NoError(t, err)
}

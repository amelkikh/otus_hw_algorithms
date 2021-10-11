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

type BitKnight struct{}

const (
	KnightNoA  = 0xfefefefefefefefe
	KnightNoH  = 0x7f7f7f7f7f7f7f7f
	KnightNoAB = 0xfcfcfcfcfcfcfcfc
	KnightNoGH = 0x3f3f3f3f3f3f3f3f
)

func (q *BitKnight) Run(strings []string) string {
	p, err := strconv.Atoi(strings[0])
	if err != nil {
		log.Fatal(fmt.Errorf("parse position: %w", err))
	}
	var K uint64
	K = 1 << p
	M := KnightNoGH&(K<<6|K>>10) |
		KnightNoH&(K<<15|K>>17) |
		KnightNoA&(K<<17|K>>15) |
		KnightNoAB&(K<<10|K>>6)

	pop := bits.OnesCount64(M)
	return fmt.Sprintf("%d\n%d", pop, M)
}

func TestBitKnight(t *testing.T) {
	dir, err := os.Getwd()
	require.NoError(t, err)
	tester := tt.NewTester(&BitKnight{}, path.Join(dir, "tests", "2.Bitboard - Конь"))
	err = tester.RunTests()
	require.NoError(t, err)
}

package util

import (
	"math"
	"math/big"
	"sort"
	"strings"

	"github.com/google/uuid"
)

// DefaultAlphabet is the default alphabet used.
const DefaultAlphabet = "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

type alphabet struct {
	chars [57]string
	len   int64
}

// Remove duplicates and sort it to ensure reproducability.
func newAlphabet() alphabet {
	var abc []string
	m := make(map[string]bool)
	for _, char := range strings.Split(DefaultAlphabet, "") {
		if _, ok := m[char]; !ok {
			m[char] = true
			abc = append(abc, char)
		}
	}

	if len(abc) != 57 {
		panic("encoding alphabet is not 57-bytes long")
	}

	sort.Strings(abc)
	a := alphabet{
		len: int64(len(abc)),
	}
	copy(a.chars[:], abc)
	return a
}

// Shot_uuid encodes uuid.UUID into a string using the least significant bits
// (LSB) first according to the alphabet. if the most significant bits (MSB)
// are 0, the string might be shorter.
func ShotUuid() string {
	var num big.Int

	u := uuid.New()
	num.SetString(strings.Replace(u.String(), "-", "", 4), 16)

	// Calculate encoded length.
	alphabet := newAlphabet()
	factor := math.Log(float64(25)) / math.Log(float64(alphabet.len))
	length := math.Ceil(factor * float64(len(u)))

	return numToString(&num, int(length), alphabet)
}

// numToString converts a number a string using the given alpabet.
func numToString(number *big.Int, padToLen int, alphabet alphabet) string {
	var (
		out   string
		digit *big.Int
	)

	for number.Uint64() > 0 {
		number, digit = new(big.Int).DivMod(number, big.NewInt(alphabet.len), new(big.Int))
		out += alphabet.chars[digit.Int64()]
	}

	if padToLen > 0 {
		remainder := math.Max(float64(padToLen-len(out)), 0)
		out = out + strings.Repeat(alphabet.chars[0], int(remainder))
	}

	return out
}

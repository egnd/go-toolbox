package assign

import (
	"github.com/cespare/xxhash"
)

// Sticky will return same number for the 2 same val's.
func Sticky(val string, cnt uint64) uint64 {
	return xxhash.Sum64String(val) % cnt
}

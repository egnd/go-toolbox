// Package assign contains hashing functions
package assign

import "math/rand"

// Random will return random number from [0...cnt) interval.
func Random(_ string, cnt uint64) uint64 {
	return uint64(rand.Intn(int(cnt))) //nolint:gosec
}

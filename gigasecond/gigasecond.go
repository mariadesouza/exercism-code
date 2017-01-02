// Package gigasecond.
package gigasecond

// Uses Package time that provides functionality for measuring and displaying time
import (
	"time"
)

// Constant declaration.
const testVersion = 4

//A gigasecond is 10^9 (1,000,000,000) seconds.
const giga = 1000000000

// AddGigasecond : Calculates the moment when someone has lived for 10^9 seconds.
// API function.  It uses a type from the Go standard library.
func AddGigasecond(t time.Time) time.Time {

	tt, _ := time.Parse("2006-01-02T15:04:05", t.Add(giga*time.Second).Format(fmtDT))
	return tt
}

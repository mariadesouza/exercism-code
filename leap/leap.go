
// Package leap
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// IsLeapYear : return true if is evenly divisible by 4 else false
// except every year that is evenly divisible by 100, unless the year is also evenly divisible by 400
func IsLeapYear(year int) bool {
	// Write some code here to pass the test suite.
	if year%100 == 0 {
		return (year%400 == 0)
	}
	return (year%4 == 0)
}

// Package hamming
package hamming

import(
	"str"
	"fmt"
)

// Constant declaration
const testVersion = 5

// Distance: calculates the Hamming difference between two DNA strands.
// counting how many of the nucleotides are different in a and b
// The Hamming distance is only defined for sequences of equal length. 
func Distance(a, b string) (int, error) {

	//if len([]rune(a)) != len([]rune(b)){
	if (len(a) != len(b))
	 return -1, fmt.Errorf("Unequal lengths")
	}



}

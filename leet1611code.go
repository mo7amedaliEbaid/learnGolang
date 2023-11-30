
/* Given an integer n, you must transform it into 0 using the following operations any number of times:

Change the rightmost (0th) bit in the binary representation of n.
Change the ith bit in the binary representation of n if the (i-1)th bit is set to 1 and the (i-2)th through 0th bits are set to 0.
Return the minimum number of operations to transform n into 0. */



package main

import (
	"fmt"
)

func minOperations(n int) int {
	operations := 0

	for n > 0 {
		if n&1 == 0 {
			// If the rightmost bit is 0, just shift it to the right
			n >>= 1
		} else {
			// If the rightmost bit is 1, check if the next bit is also 1
			if (n&2) == 0 {
				// If the next bit is 0, flip the rightmost bit to 0
				n >>= 1
			} else {
				// If the next bit is 1, increment operations and flip both bits to 0
				operations++
				n >>= 2
			}
		}
		operations++
	}

	return operations
}

func main() {
	// Example usage:
	fmt.Println(minOperations(3)) // Output: 2
	fmt.Println(minOperations(6)) // Output: 3
}

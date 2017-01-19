// Dyson Doubler by Johan Dahl

package main

import (
	"fmt"
	"os"
)

func main() {
	var i uint64 = 1
	var e uint64 = 1
	var lim uint64 = 10
	// Eg if i = 123 is e = 100 and lim = 1000
	// the last = 3 and d = 12 + 3 * 100 = 312

	for {
		fmt.Println(i)
		last := i % 10 // Get last digit
		if last != 0 {
			d := i/10 + last*e // but last + last * e
			// If d is double of i are the ready
			if i*2 == d {
				fmt.Println("FOUND")

				file, err := os.Create("/tmp/result.txt")
				if err != nil {
					panic(err)
				}
				fmt.Fprintln(file, i)
				defer file.Close()
				return
			}
		}
		i++
		// If n get over lim make e and lim 10 times bigger
		if i >= lim {
			lim *= 10
			e *= 10
		}
	}
}
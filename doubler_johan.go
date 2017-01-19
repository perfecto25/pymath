// Dyson Doubler by Johan Dahl

package main

import (
	"fmt"
	"os"
)

/*
Start with i = 11 and d(ouble) = d0 = 11. And for each numbers 1 .. 9
Increase i + 1 and d + 10
Then does d start with d0 + 1
If i > limit (100 first time) is
	d0 = 11 * 10
	limit = limit * 10



11 => 11
12 => 21 (11 + 10)
13 => 31 (21 + 10)
20 .. just ignore
21 => 12

100 ... just ignore
101 => 110 (11 * 10)
102 => 210 (110 + 100)

and so on...

This is about 2.5 faster than previous because
we don't do alot of multiplication and modulus each round
just adding and one multiplication

*/
func main() {
	var i uint64 = 11
	var d0 uint64 = 11
	var e uint64 = 10
	var lim uint64 = 100

	for {

		d := d0
		// Jump over numbers ending in zero
		for n := 1; n <= 9; n++ {
			fmt.Println(i, d) // Remove this line for even faster running
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
			i++
			d += e
		}
		i++
		d0++
		if i >= lim {
			d0 = 11 * e
			e *= 10
			lim *= 10
		}
	}
}

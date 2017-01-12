package main

import (
	"fmt"
	"os"
	"strconv"
)

func rotate(in uint64) (out uint64) {

	last := in % 10
	rest := in / 10

	var buf []byte
	buf = strconv.AppendUint(buf, last, 10)
	buf = strconv.AppendUint(buf, rest, 10)

	out, err := strconv.ParseUint(string(buf), 10, 64)
	if err != nil {
		panic(err)
	}

	return

}

func main() {

	var i uint64 = 1

	for i = 1; ; i++ {

		fmt.Println(i)
		if 2*i == rotate(i) {
			fmt.Println("FOUND")

			file, err := os.Create("/tmp/result.txt")
			if err != nil {
				panic(err)
			}
			fmt.Fprintln(file, i)

			break
		}

	}

}

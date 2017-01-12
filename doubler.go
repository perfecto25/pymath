package main

import (
	"fmt"
	//"reflect"
	"os"
	"strconv"
	"strings"
)

func countIntegers(arr []string) int {
	return len(arr)
}

// moves last ingteger to first position in an array
func moveLastToFirst(number uint64) uint64 {
	str := strconv.Itoa(number)

	// string array to store each number
	arr := []string{}

	// loop thru string, add each number to slice
	for _, x := range str {
		number := string(x)
		arr = append(arr, number)
	}

	fmt.Println(countIntegers(arr))
	fmt.Println("\n")

	// get last # of array
	lastNum := arr[len(arr)-1]

	// move last number to 1st position
	arr = append([]string{lastNum}, arr...)

	// get index of last element
	lastindex := len(arr) - 1

	// remove last number from array (duplicate)
	arr = append(arr[:lastindex])

	// join array elements into single number
	arrString := strings.Join(arr[:], "")

	// convert string array to integer
	num, _ := strconv.Atoi(arrString)

	return num
}

func compare(i uint64) string {

	double := i * 2
	ltf := moveLastToFirst(i)

	if double == ltf {

		return "found"
	}

	return "continue"
}

func main() {

	var i uint64 = 1

	for {

		fmt.Println(i)
		if compare(i) == "found" {
			fmt.Println("FOUND")

			file, err := os.Create("/tmp/result.txt")
			if err != nil {
				panic(err)
			}

			fmt.Fprintln(file, i)

			defer file.Close()

			break

		} else {
			i += 1
		}
	}

	// fmt.Println(arrString)
	// fmt.Println(reflect.TypeOf(num)) // get var type

	//	num := strconv.Atoi(arrString)
	//
	// fmt.Println("last number", numSlice[lastNum])

	// fmt.Println(numSlice)
	// lastNum := numSlice[len(numSlice)-1]
	// numSlice[0] = lastNum
	// numSlice = append(numSlice, lastNum)
	// fmt.Println(lastNum)
	// fmt.Println(numSlice)

	//strArr := strings.Fields(str)

}

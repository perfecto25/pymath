package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// func moveLastToFirst(num int) {

// 	str := strconv.Itoa(num)
// 	strArray := strings.Fields(str)
// 	return strArray[1:2]
// }

func main() {
	// somevar := moveLastToFirst(345)
	// fmt.Println(somevar)
	str := strconv.Itoa(123)

	// string slice to store each number
	numSlice := make([]string, 0)

	// loop thru string, add each number to slice
	for _, x := range str {
		number := string(x)
		numSlice = append(numSlice, number)
	}

	fmt.Println(numSlice)
	lastNum := numSlice[len(numSlice)-1]
	fmt.Println(lastNum)

	strArr := strings.Fields(str)
	fmt.Println(reflect.TypeOf(strArr))

}

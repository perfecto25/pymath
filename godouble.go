package main

import (
	"fmt"
	//"reflect"
	"strconv"
	"strings"
)

func moveLastToFirst(num int) []string {

	str := strconv.Itoa(num)
	fmt.Println(str)
	strArray := strings.Fields(str)
	return strArray

}

func main() {
	res := moveLastToFirst(305941)
	//fmt.Println(reflect.TypeOf(res))
	fmt.Println(res[0])

}

// string = str(num)

// mylist = list(string)

// # for position, item in enumerate(mylist):
// #     print "position: %s, item: %s" % (position, item)

// last = mylist[-1]

// length = len(mylist)

// mylist.insert(0,last)

// last = mylist[-1]

// length = len(mylist)

// del mylist[length-1]

// mylist = ''.join(mylist)

// num = int(mylist)

// return num

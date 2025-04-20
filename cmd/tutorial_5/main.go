package main

import (
	"fmt"
	"strings"
)

func main() {

	/*var myString = "résume"
	var indexed = myString[0]
	fmt.Printf("%v, %T \n", indexed, indexed) // prints 114

	for i, v := range myString {
		fmt.Println(i, v)
	}*/

	var myString = []rune("résume")
	var indexed = myString[1]
	fmt.Printf("%v, %T \n", indexed, indexed) // prints 114

	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("\n the length of 'myString' is %v", len(myString))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v \n", myRune)

	var strSlice = []string{"a", "b", "c"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v\n", catStr)
	
	// inefficient way of building a string
	/*var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Println(catStr)*/

}

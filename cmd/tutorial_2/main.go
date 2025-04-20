package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	fmt.Println(intNum)

	var floatNum float32 = 12345.9
	var intNum32 int32 = 2
	var result = floatNum + float32(intNum32)
	fmt.Println(result)

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	fmt.Println(len(myString))
	fmt.Println(len("A"))
	fmt.Println(utf8.RuneCountInString("y"))

	var myRune rune = 'A'
	fmt.Println(myRune) //asci code

	var myBoolean = true
	fmt.Println(myBoolean)

	var intNum3 int
	fmt.Println(intNum3) // default value of int is 0

	var guessType = "hey" // guess the type of a variable by its value
	fmt.Println(guessType)

	guessTypeAndVar := "stringing"
	fmt.Println(guessTypeAndVar) // short variable declaration

	var1, var2 := 1, 2
	fmt.Println(var1, var2)


	const myConst string = "cannot change this "
	fmt.Println(myConst)
	/*errors :
		const myConst string      //you need to give it a value
		you cannot change the value of a const
	
	*/
}

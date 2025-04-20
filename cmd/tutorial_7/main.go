package main

import "fmt"

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\n the memory location of variable in square function :  %p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

//pointers edition :
func square2(thing2 *[5]float64) {
	fmt.Printf("\n the memory location of variable in square function :  %p", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
}

func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("the value p points is : %v", *p)
	fmt.Printf("\n the value if i is  %v \n", i)

	i = 10
	p = &i
	fmt.Printf("the value p points is : %v", *p)
	fmt.Printf("\nadress of i : %v , adress that p points to : %v \n", p, &i)

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice // under the hood slices contains pointers to underline array
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\n the memory location of variable in main function :  %p", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("\n the result is : %v", result)
	fmt.Println(thing1)

	square2(&thing1)
	fmt.Printf("\n the result is : %v", thing1)
	// just use pointers like in C

}

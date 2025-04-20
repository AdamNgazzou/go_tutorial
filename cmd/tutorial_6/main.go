package main

import "fmt"

//just like in C language
type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner   // or ownerInfo owner
	int
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

//method for electric engine (function)
func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type owner struct {
	name string
}

//interfaces (global function for any engine):
type engine interface {
	//method signature
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it ")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}, 5} // or gasEngine{mpg: 25, gallons: 15, owner: owner{"Alex"}}  //default 0 0
	myEngine.mpg = 5
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name, myEngine.int)

	fmt.Printf("total miles left in tank %v \n", myEngine.milesLeft())
	canMakeIt(myEngine, 50)
	canMakeIt(myEngine, 80)


}

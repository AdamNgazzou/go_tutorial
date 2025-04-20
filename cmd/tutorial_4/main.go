package main

import "fmt"

func main() {
	var intArr [3]int32
	intArr[1] = 1
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intArr2 = [...]int32{1, 2, 3} //the length emited by the compiler by the [...]
	fmt.Println(intArr2)

	//slices
	var slice []int32 = []int32{1, 2, 3}
	fmt.Println(slice)
	fmt.Printf("the length is %v with capacity %v", len(slice), cap(slice))

	//unlike arrays i can do this with slices :
	slice = append(slice, 7)
	fmt.Printf("\n the length is %v with capacity %v", len(slice), cap(slice))
	fmt.Println(slice)

	var slice2 []int32 = []int32{8, 9}
	slice = append(slice, slice2...)
	fmt.Println(slice)

	var slice3 []int32 = make([]int32, 3, 8)
	fmt.Println(slice3)

	//maps
	var Mymap map[string]uint8 = make(map[string]uint8)
	fmt.Println(Mymap)

	var Mymap2 = map[string]uint8{"Adam": 23, "Nadine": 50}
	fmt.Println(Mymap2)
	fmt.Println(Mymap2["Adam"])
	fmt.Println(Mymap2["doesn't exist name"])
	var age, ok = Mymap2["a name that doesn't Exist"]
	if ok {
		fmt.Printf("here is the age %v", age)
	} else {
		fmt.Println("well well well it doesn't exist")
	}
	//delete an element from a map :
	delete(Mymap2, "Nadine")
	fmt.Println(Mymap2)

	//loops  map:
	for name, age := range Mymap2 {
		fmt.Printf("Name : %v and age : %v\n", name, age)
	}

	//loops array :
	for i, v := range intArr {
		fmt.Printf("Index: %v , Value! %v \n", i, v)
	}

	// go doesn't have a while loop we can achive that tho with a for loop that has a break :
	var i int = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}
	// easier way to write it (just like c language )
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}
}

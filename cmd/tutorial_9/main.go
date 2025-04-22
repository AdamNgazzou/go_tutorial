package main

// channels : a way to make goroutines to pass around information(data)
// 1/ they hold data
// 2/ they are thread safe
// 3/ listen for data
//buffer channel = storing multiple values in the channel
import (
	"fmt"
	"math/rand"
	"time"
)

/*func main() {
	var c = make(chan int, 5)
	go process(c)
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}

}

func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("exiting process")
}*/
//realistic :

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE = 3

func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)

	var websites = []string{"walmart.com", "costco.com", "wholefood.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)

}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			tofuChannel <- website
			break
		}
	}

}
func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}

}
func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select { // if statement for channels
	case website := <-chickenChannel:
		fmt.Printf("\nFound a deal on chicken at %v", website)
	case website := <-tofuChannel:
		fmt.Printf("\nFound a deal on chicken at %v", website)
	}
}
//need to be back to this and routies

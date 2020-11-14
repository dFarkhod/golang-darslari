// Channel'lar haqida

package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for {
			channel1 <- "Tez"
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for {
			channel2 <- "Sekin"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		fmt.Println(<-channel1)
		fmt.Println(<-channel2)
	}
}

/* // 2-misol
func main() {
	channel := make(chan int)
	go getRandomNumber(channel)
	for randomNumber := range channel {
		fmt.Println("tasodifiy son: ", randomNumber)
	}

}

func getRandomNumber(channel chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 3; i++ {
		number := rand.Intn(1000)
		time.Sleep(time.Second * 1)
		channel <- number
	}
	close(channel)
}
*/

/* // 1-misol
func main() {

	channel := make(chan string)
	go func() {
		channel <- "anor"
		channel <- "olma"
	}()
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
*/

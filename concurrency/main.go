package main

import (
	"fmt"
	"time"
)

func main() {
	greet("I am a synchronous function")

	go slowGreet("I am a slow goroutine")
	go greet("I am a goroutine")
	fmt.Println("Main function")

}

func greet(phrase string) {
	fmt.Println("Hello, World!", phrase)
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello, World!", phrase)
}

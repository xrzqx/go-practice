package main

import (
	"fmt"
	"strconv"
)

func RangeChannel()  {
	channel := make(chan string)
	go func ()  {
		for i := 0; i < 10; i++ {
			channel <- "Iterator - " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel{
		fmt.Println("Recive ",data)
	}

}

func main() {
	RangeChannel()
}
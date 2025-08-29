package belajargolanggoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func(){
		channel <- "Ryan Fajri"
		time.Sleep(2 * time.Second)
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <- channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	channel <- "Ryan Fajri"
	time.Sleep(2 * time.Second)
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn (channel chan<- string) {
	channel <- "Ryan Fajri"
	time.Sleep(2 * time.Second)
}

func OnlyOut (channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Ryan"
	channel <- "Fajri"
	channel <- "Alifah"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("Selesai")
}


func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func(){
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}

		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select{
			case data := <-channel1:
				fmt.Println("Data dari channel 1", data)
				counter++
			case data := <-channel2:
				fmt.Println("Data dari channel 2", data)
				counter++
			default:
				fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}

	}

	
}
package main

import (
	"felixfern/go-hello/area"
	"felixfern/go-hello/concurrency"
	"felixfern/go-hello/dictionary"
	"felixfern/go-hello/wallet"
	"fmt"
	"time"
)

func mockWebsiteChecker(url string) bool {
	time.Sleep(100 * time.Millisecond)

	return url != "waat://furhurterwe.geds"
}

func main() {
	test := []struct {
		name  string
		shape area.Shape
	}{

		{name: "Rectangle", shape: area.Rectangle{Width: 10.0, Height: 20.0}},
		{name: "Circle", shape: area.Circle{Radius: 2.5}},
	}

	wallets := []wallet.Wallet{
		{Id: 1, Name: "John", Balance: 12000.50},
		{Id: 2, Name: "Jane", Balance: 50000},
	}

	for _, val := range test {
		fmt.Println(fmt.Sprintf("%s %f", val.name, val.shape.Area()))
	}

	wallets[0].Deposit(5000)
	wallets[1].Deduct(5000)

	fmt.Println(wallets[0].GetBalance())
	fmt.Println(wallets[1].GetBalance())

	dict := dictionary.Dictionary{
		"saya": {En: "Me", Id: "saya"},
	}

	func() {
		res, err := dict.Search("saya")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(res)
	}()

	dict.Add("coba", dictionary.Language{En: "Try", Id: "coba"})
	dict.Update("saya", dictionary.Language{En: "Update", Id: "Update Id"})

	func() {
		res, err := dict.Search("saya")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(res)
	}()

	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	concurrency.CheckWebsites(mockWebsiteChecker, websites)
	concurrency.CheckWebsitesConcurrent(mockWebsiteChecker, websites)
}

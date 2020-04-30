package main

import (
	"fmt"
	"time"

	"github.com/mattwangwl/txdis"
)

var appleStorage = 100
var bananaStorage = 80

func main() {
	d := txdis.New(txdis.NewDefaultConfig())

	go func() {
		for i := 0; i < 50; i++ {
			apple := txdis.NewTask(1, func() (interface{}, error) {
				appleStorage = appleStorage - 1
				return nil, nil
			})

			go d.Delegate(apple)
		}
	}()

	go func() {
		for i := 0; i < 60; i++ {
			banana := txdis.NewTask(2, func() (interface{}, error) {
				bananaStorage = bananaStorage - 1
				return nil, nil
			})

			go d.Delegate(banana)
		}
	}()

	<-time.After(5 * time.Second)
	fmt.Println("appleStorage:", appleStorage)
	fmt.Println("bananaStorage:", bananaStorage)
	// Output:
	// appleStorage: 50
	// bananaStorage: 20
}

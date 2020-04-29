package main

import (
	"fmt"

	"github.com/mattwangwl/txdis"
)

func main() {
	hello := txdis.NewTask(123, func() (interface{}, error) {
		return "hello", nil
	})

	world := txdis.NewTaskR(func() (interface{}, error) {
		return "world", nil
	})

	d := txdis.New(txdis.NewDefaultConfig())
	d.Delegate(world)
	d.Delegate(hello)

	fmt.Println(hello.Result())
	fmt.Println(world.Result())
	// Output:
	// {hello <nil>}
	// {world <nil>}
}

package main

import (
	"fmt"
	"time"

	"github.com/mattwangwl/txdis"
)

func main() {
	d := txdis.New(txdis.NewDefaultConfig())

	task1 := txdis.NewTask(123, func() (interface{}, error) {
		for i := 0; i < 5; i++ {
			fmt.Println("task1 num:", i)
			<-time.After(100 * time.Millisecond)
		}
		fmt.Println("task1 finish")
		return "task1 result", nil
	})

	task2 := txdis.NewTask(456, func() (interface{}, error) {
		for i := 0; i < 3; i++ {
			fmt.Println("task2 num:", i)
			<-time.After(100 * time.Millisecond)
		}
		fmt.Println("task2 finish")
		return "task2 result", nil
	})

	d.Delegate(task1)
	d.Delegate(task2)

	fmt.Println(task1.Result())
	fmt.Println(task2.Result())
	// Output:
	// task2 num: 0
	// task1 num: 0
	// task2 num: 1
	// task1 num: 1
	// task1 num: 2
	// task2 num: 2
	// task2 finish
	// task1 num: 3
	// task1 num: 4
	// task1 finish
	// {task1 result <nil>}
	// {task2 result <nil>}
}

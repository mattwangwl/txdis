package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/mattwangwl/txdis"
)

func main() {
	d := txdis.New(txdis.NewDefaultConfig())

	go func() {
		for {
			d.Delegate(txdis.NewTaskR(func() (interface{}, error) {
				<-time.After(100 * time.Millisecond)
				return "result", nil
			}))
			<-time.After(100 * time.Millisecond)
		}
	}()

	http.ListenAndServe("0.0.0.0:8080", nil)
}

package txdis

import (
	"context"
	"testing"
	"time"
)

// test runtime error: index out of range
// https://github.com/golang/go/issues/3611
func Test_random_Int64(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	r := newRandom(time.Now().Unix())
	tk := time.NewTicker(2 * time.Second)

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 50; i++ {
		go func(j int) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					r.Int63()
				}

			}
		}(i)
	}

	select {
	case <-tk.C:
		cancel()
	}
}

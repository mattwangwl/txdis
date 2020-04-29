package txdis

import (
	"math/rand"
	"sync"
	"time"
)

var _random = newRandom(time.Now().Unix())

func newRandom(seed int64) *random {
	r := &random{}
	source := rand.NewSource(seed)
	r.generator = rand.New(source)
	return r
}

type random struct {
	mx        sync.Mutex
	generator *rand.Rand
}

func (r *random) Int63() int64 {
	r.mx.Lock()
	i := r.generator.Int63()
	r.mx.Unlock()
	return i
}

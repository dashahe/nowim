package ulid

import (
	"github.com/oklog/ulid/v2"
	"math/rand"
	"time"
)

var ulids chan string

func init() {
	ulids = make(chan string, 1000)
	go func() {
		entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
		for {
			id := ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
			ulids <- id
		}
	}()
}

func ULID() string {
	return <-ulids
}

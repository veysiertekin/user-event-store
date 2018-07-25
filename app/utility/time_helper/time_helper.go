package time_helper

import (
	"time"
	"math/rand"
)

func ToMilliSeconds(d time.Duration) float64 {
	msec := d / time.Millisecond
	nsec := d % time.Millisecond
	return float64(msec) + float64(nsec)/(1e7)
}

func WaitForARandomDuration(maxWaitDurationInMilliseconds int) {
	amt := time.Duration(rand.Intn(maxWaitDurationInMilliseconds))
	time.Sleep(time.Millisecond * amt)
}

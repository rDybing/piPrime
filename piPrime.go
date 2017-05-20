package main

import (
	"fmt"
	"math"
	"time"
)

type output_t struct {
	number uint64
	time   float32
	prime  int32
}

type timer_t struct {
	timeOld int64
	timeNew int64
}

func main() {
	// manually setting CPU cores do nothing after go1.5
	// runtime.GOMAXPROCS(4)
	var countTo uint64
	countTo = 1600000
	fmt.Println("here we go. Ol'skool serial prime-finder")
	primeSequential(countTo)
	fmt.Println("and now the concurrent way!")
	primeConcurrent(countTo)
}

func primeSequential(limit uint64) {
	var out output_t
	var timer timer_t
	timer.timeOld = millis()

	for out.number = 2; out.number < limit; out.number++ {
		if isPrime(out.number) {
			out.prime++
		}
		if (out.number)%1000000 == 0 {
			out.time = getTimePassed(&timer)
			output(&out, 0)
		}
	}
	out.time = getTimePassed(&timer)
	output(&out, 0)
}

func primeConcurrent(limit uint64) {
	var out output_t
	var timer timer_t
	var isTrue bool

	chRec := make(chan bool, 100)
	chSend := make(chan uint64, 100)

	timer.timeOld = millis()

	// spin up four goroutines - one for each CPU core
	for i := 0; i < 4; i++ {
		go isPrimeConc(chSend, chRec)
	}

	for out.number = 2; out.number < limit; out.number += 4 {
		// send
		for i := 0; i < 4; i++ {
			chSend <- out.number + uint64(i)
		}
		// receive
		for i := 0; i < 4; i++ {
			isTrue = <-chRec
			if isTrue {
				out.prime++
			}
			if (out.number+uint64(i))%1000000 == 0 {
				out.time = getTimePassed(&timer)
				output(&out, i)
			}
		}
	}
	out.time = getTimePassed(&timer)
	output(&out, 0)
}

func isPrimeConc(inC <-chan uint64, out chan<- bool) {
	var i uint64
	var s uint64
	var in uint64
	var flag bool
	var exit bool

	for {
		exit = false
		in = <-inC

		s = uint64(math.Sqrt(float64(in)))

		if in > 2 && in%2 == 0 {
			flag = false
			exit = true
		}
		if exit == false {
			for i = 3; i <= s; i += 2 {
				if in%i == 0 {
					flag = false
					exit = true
					break
				}
			}
		}
		if exit == false {
			flag = true
		}
		out <- flag
	}
}

func isPrime(number uint64) bool {
	var i uint64
	var s uint64
	s = uint64(math.Sqrt(float64(number)))
	if number > 2 && number%2 == 0 {
		return false
	}
	for i = 3; i <= s; i += 2 {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func output(o *output_t, i int) {
	fmt.Printf("%8d || time: %6.2f || primes: %d\n", o.number+uint64(i), o.time, o.prime)
}

func getTimePassed(t *timer_t) float32 {
	var timeSec float32
	var timeTemp int64
	t.timeNew = millis()
	timeTemp = t.timeNew - t.timeOld
	timeSec = float32(timeTemp) / float32(1000)
	return timeSec
}

func millis() int64 {
	return time.Now().Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

package main

import (
	"fmt"
	"math"
	"time"
)

type output_t struct {
	number uint64
	time   time.Duration
	prime  int32
}

func main() {
	// manually setting CPU cores do nothing after go1.5
	// runtime.GOMAXPROCS(4)
	var countTo uint64
	countTo = 16000000
	fmt.Println("here we go. Ol'skool serial prime-finder")
	primeSequential(countTo)
	fmt.Println("and now the concurrent way!")
	primeConcurrent(countTo)
}

func primeSequential(limit uint64) {
	var out output_t
	now := time.Now()

	for out.number = 2; out.number < limit; out.number++ {
		if isPrime(out.number) {
			out.prime++
		}
		if (out.number)%1000000 == 0 {
			out.time = time.Since(now)
			output(&out, 0)
		}
	}
	out.time = time.Since(now)
	output(&out, 0)
}

func primeConcurrent(limit uint64) {
	var out output_t

	chRec := make(chan bool, 100)
	chSend := make(chan uint64, 100)

	now := time.Now()

	// spin up four goroutines - one for each CPU core
	for i := 0; i < 4; i++ {
		go isPrimeConc(chSend, chRec)
	}
	go func(limit uint64, chSend chan uint64) {
		for i := uint64(2); i < limit; i++ {
			// send
			chSend <- i
		}
	}(limit, chSend)
	// receive
	for i := uint64(2); i < limit; i++ {
		if <-chRec {
			out.prime++
		}
		if i%1000000 == 0 {
			out.time = time.Since(now)
			output(&out, i)
		}
	}
	out.number = limit
	out.time = time.Since(now)
	output(&out, 0)
	close(chRec)
	close(chSend)
}

func isPrimeConc(inC <-chan uint64, out chan<- bool) {
	for i := range inC {
		out <- isPrime(i)
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

func output(o *output_t, i uint64) {
	fmt.Printf("%8d || time: %v || primes: %d\n", o.number+i, o.time, o.prime)
}

// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
	"math/rand"
)

type NumberPredicate interface {
	Fulfill(n int) bool
}

type Predicate func(n int) bool

func (p Predicate) Fulfill(n int) bool {
	return p(n)
}

type IntSequence interface {
	Next() int
}

type Sequence struct {
	current         int
	numberPredicate NumberPredicate
}

func (s *Sequence) Next() int {
	var next int
	for n := s.current + 1; n <= math.MaxInt32; n++ {
		if s.numberPredicate.Fulfill(n) {
			next = n
			break
		}
	}
	s.current = next
	return next
}

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isInteger(n int) bool {
	return true
}

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	var intSequence IntSequence

	random := rand.Intn(3)
	switch random {
	case 0:
		intSequence = &Sequence{current: 0, numberPredicate: Predicate(isPrime)}
	case 1:
		intSequence = &Sequence{current: 0, numberPredicate: Predicate(isInteger)}
	case 2:
		intSequence = &Sequence{current: 0, numberPredicate: Predicate(isEven)}
	}

	for i := 1; i <= 26; i++ {
		fmt.Println(intSequence.Next())
	}

}

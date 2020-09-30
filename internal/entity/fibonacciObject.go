package entity

import (
	"time"
)

/**
The basic object which stores the nth fibonacci number for any given number,
however, it is assumed the given number input would be a natural number.
*/
type FibonacciObject struct {
	ID                 string    `json:"id"`
	NthInteger         uint64    `json:"inputNum"`
	NthFibonacciNumber uint64    `json:"fiboNum"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

/**
Some method of the object to change the fields
*/
func (f *FibonacciObject) ModifyFirstVal(newID string) {
	f.ID = newID
}
func (f *FibonacciObject) ModifySecondVal(newnthInteger uint64) {
	f.NthInteger = newnthInteger
}
func (f *FibonacciObject) ModifyThirdVal(newnthFibonacciNumber uint64) {
	f.NthFibonacciNumber = newnthFibonacciNumber
}
func (f *FibonacciObject) ModifyFourthVal() {
	f.CreatedAt = time.Now()
}

func (f *FibonacciObject) ModifyFifthVal() {
	f.UpdatedAt = time.Now()
}

/**
HOW TO FIND OUT THE NTH FIBONACCI NUMBER???
SOLUTION:
Recursive approach with Dynamic Programming
Big O complexity for the implemented methods will be
		Time Complexity: O(n)
		Space Complexity: O(n)
*/
func (f *FibonacciObject) FindFibonacci(nthNumber uint64) uint64 {
	// memoization to store key value pairs temporarily
	mem := make(map[uint64]uint64)
	fiboFounded := fib(nthNumber, &mem)

	return fiboFounded
}

/**
A recursive function which calculates the coresponding Fibonacci number
*/
func fib(nthNumber uint64, mem *map[uint64]uint64) uint64 {
	if nthNumber == 0 {
		return 0
	} else if nthNumber <= 2 {
		return 1
	} else if val, ok := (*mem)[nthNumber]; ok {
		return val
	}
	val := fib(nthNumber-1, mem) + fib(nthNumber-2, mem)
	(*mem)[nthNumber] = val
	return val
}

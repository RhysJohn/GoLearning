package main

import (
	"fmt"
	"strconv"

	"codewars/sumofperfectsquares"
)

func main() {
    tests := []uint64 { 661915703, 999887641, 999950886 };
    for _, test := range tests {
        list := sumofperfectsquares.SumOfSquares(test);
        fmt.Println( "For number: " + strconv.Itoa(int(test)) + " Answer: " + strconv.Itoa(int(list))) 
    }
}
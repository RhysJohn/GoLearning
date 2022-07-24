// Task Description: Given an integer n (3 < n < 10^9), find the length of the smallest list of perfect squares which add up to n.
// Come up with the best alogrithm you can; you'll need it.

// Example - sum_of_squares(17) = 2
// 17 = 16 + 1

// 1. Create a fucntion that takes in a number.
// 2. Find out the biggest square that can fit within it.
// 3. Recursion with the remainder and return.

package sumofperfectsquares

import (
	"fmt"
	"strconv"
)

// First attempt.
// func SumOfSquares(n uint64) uint64 {
// 	fmt.Println("n is " + strconv.Itoa(int(n)));
// 	if n < 4 {
// 		/* Because we are lower than 2 squared whatever value n holds this
// 		   is the amount of 1 squared are going to be added to the list.
// 		*/
// 		return n;
// 	}

// 	// Go through from [1,2] squared to see if the sum of these are above n
// 	var index uint64;
// 	index = 1;
// 	var result uint64;
// 	result = 0;
// 	for {
// 		// fmt.Println("Index is " + strconv.i(index));

// 		var squaresSum = findSquaresSum(index, index + 1);
// 		fmt.Println("Checking n against sum of squares " + strconv.Itoa(int(squaresSum)));

// 		if squaresSum >= n {
// 			var remainder = n - (index + 1) * (index + 1);
// 			fmt.Println("Remainder = " + strconv.Itoa(int(remainder)));
// 			result++;
// 			result += SumOfSquares(remainder);
// 			break;
// 		} else {
// 			index++;
// 		}
// 	}

// 	fmt.Println("Result = " + strconv.Itoa(int(result)));

// 	return result;
// }

type keyPairValue struct {
	key uint64
	value uint64
}

type solution struct {
	length uint64
	perfectSquares []uint64
}

// Second Attempt.
func SumOfSquares(n uint64) int {
	// Declare an array and store all the possible squares which fit into n.
	var squareValuesWhichFit []keyPairValue; 

	// This contains all the solution for the given lengths.
	var answers []solution;

	var index uint64 = 1;
	for {
		var indexSquared = index * index;
		if n >= uint64(indexSquared) {
			possibleValue := keyPairValue {
				key: index,
				value: indexSquared,
			}
			fmt.Println("Possible value: [" + strconv.Itoa(int(possibleValue.key)) + ", " +  strconv.Itoa(int(possibleValue.value)) + "]");

			squareValuesWhichFit = append(squareValuesWhichFit, possibleValue);
			continue;
		} else {
			// Square of the index > n, therefore we have all possible values. 
			break;
		}
	}

	// Based on length of squareValuesWhichFit, we find the minimun per length and then reduce
	var length = len(squareValuesWhichFit);
	for i := length; i > 0; i-- {		
		var possibleValues []uint64;
		for j := 0; j < i; j++ {
			possibleValues = append(possibleValues, squareValuesWhichFit[i].value);
		}
		
		var lengthSolution = findSolution(n, possibleValues);
		lengthAnswer := solution {
			length: uint64(i),
			perfectSquares: lengthSolution,
		};

		answers = append(answers, lengthAnswer);
	}

	
	// Find the element with the smallest length in answers.
	return len(answers);
}

func findSolution(n uint64, valuesToPlayWith []uint64) []uint64 {
	var solution []uint64;
	for i := len(valuesToPlayWith); i > 0; i++ {
		var element uint64 = valuesToPlayWith[i];
		if n > element {
						
		}
	}

	if squaresSum >= n {
		var remainder = n - (index + 1) * (index + 1);
		fmt.Println("Remainder = " + strconv.Itoa(int(remainder)));
		result++;
		result += SumOfSquares(remainder);
		break;
	} else {
		index++;
	}
}

func findSquaresSum(a uint64, b uint64) uint64 {
	var aSqaured = a * b;
	var bSquared = b * b;

	var sum = aSqaured + bSquared;
	fmt.Println(sum)

	return sum;
}
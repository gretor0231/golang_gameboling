package functions

import "fmt"

//Complete the function f that sorts an integer array by ascending order. For example, given array 2 1, function f should output 1 2.

func Two(input []int) string {
	//sort the array in ascending order
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
	//return the sorted array as a string
	return fmt.Sprint(input)
}

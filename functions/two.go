package functions

import "fmt"

//sort the array in ascending order
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

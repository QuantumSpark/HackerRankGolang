package main

import (
	"bufio"
	"fmt"
	"os"
)

func roundGrade(arr []int) []int {
	l := len(arr)
	roundedGrade := make([]int, l)

	for i := 0; i < l; i++ {
		var currentGrade = arr[i]

		if currentGrade < 38 {
			roundedGrade[i] = currentGrade
		} else {
			var nextMultiple5 = (currentGrade + 4) / 5 * 5
			if nextMultiple5-currentGrade < 3 {
				roundedGrade[i] = nextMultiple5
			} else {
				roundedGrade[i] = currentGrade
			}
		}

	}

	return roundedGrade
}

func main() {
	var n int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(io, &arr[i])
	}
	var roundedGrade = roundGrade(arr)
	for i := 0; i < n; i++ {
		fmt.Println(roundedGrade[i])
	}
}

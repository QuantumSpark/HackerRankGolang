package main

import (
	"bufio"
	"fmt"
	"os"
)

func uniqueString(s string) bool {

	arr := make([]int, 256)

	for _, r := range s {
		if arr[r] == 0 {
			arr[r] = 1
		} else if arr[r] == 1 {
			return false
		}
	}
	return true
}

func main() {

	io := bufio.NewReader(os.Stdin)
	var s string

	fmt.Fscan(io, &s)

	fmt.Println(uniqueString(s))

}

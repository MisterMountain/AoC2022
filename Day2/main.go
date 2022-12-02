package main

import (
	"fmt"
	"os"
)

func main() {
	var result1 int
	var result2 int
	var data, _ = os.Open("puzzle")
	var buffer = make([]byte, 10000)
	data.Read(buffer)
	for i := 0; i < len(buffer); i += 4 {
		if buffer[i] == 65 {
			if buffer[i+2] == 88 {
				result1 = result1 + 3
				result2 = result2 + 0 + 3
			} else if buffer[i+2] == 89 {
				result1 = result1 + 6
				result2 = result2 + 3 + 1
			} else if buffer[i+2] == 90 {
				result1 = result1 + 0
				result2 = result2 + 6 + 2
			}
		} else if buffer[i] == 66 {
			if buffer[i+2] == 88 {
				result1 = result1 + 0
				result2 = result2 + 0 + 1
			} else if buffer[i+2] == 89 {
				result1 = result1 + 3
				result2 = result2 + 3 + 2
			} else if buffer[i+2] == 90 {
				result1 = result1 + 6
				result2 = result2 + 6 + 3
			}
		} else if buffer[i] == 67 {
			if buffer[i+2] == 88 {
				result1 = result1 + 6
				result2 = result2 + 0 + 2
			} else if buffer[i+2] == 89 {
				result1 = result1 + 0
				result2 = result2 + 3 + 3
			} else if buffer[i+2] == 90 {
				result1 = result1 + 3
				result2 = result2 + 6 + 1
			}
		}
		result1 = result1 + int(buffer[i+2]-87)
	}
	fmt.Println("Result 1:", result1)
	fmt.Println("Result 2:", result2)
}

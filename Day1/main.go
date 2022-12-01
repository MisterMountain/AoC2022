package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var file, _ = os.Open("puzzle")
	defer file.Close()
	scan := bufio.NewScanner(file)

	data := make([]int, 1)
	max := 0
	for scan.Scan() {
		if line := scan.Text(); len(line) > 0 {
			c, _ := strconv.Atoi(line)
			data[len(data)-1] += c
			if data[len(data)-1] > max {
				max = data[len(data)-1]
			}
		} else {
			data = append(data, 0)
		}
	}
	fmt.Println(max)
	sort.Ints(data)
	fmt.Println(data[len(data)-3] + data[len(data)-2] + data[len(data)-1])
}

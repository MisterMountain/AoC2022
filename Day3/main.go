package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed puzzle
var input string

func main() {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	task1(lines)
	task2(lines)
}

func task1(lines []string) {
	var count int
	for _, line := range lines {
		a := line[0 : len(line)/2]
		b := line[len(line)/2:]

		for _, c := range a {
			if strings.Contains(b, string(c)) {
				p := priority(c)
				count += p
				break
			}
		}
	}
	fmt.Println(count)
}

func task2(lines []string) {
	var count int
	for i := 0; i < len(lines); i += 3 {
		for _, char := range lines[i] {
			if strings.Contains(lines[i+1], string(char)) && strings.Contains(lines[i+2], string(char)) {
				count += priority(char)
				break
			}
		}
	}
	fmt.Println(count)
}

func priority(c int32) int {
	if 'a' <= c && c <= 'z' {
		return int(1 + c - 'a')
	} else {
		return int(27 + c - 'A')
	}
}

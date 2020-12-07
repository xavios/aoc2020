package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("third.txt")
	scanner := bufio.NewScanner(f)
	var oMap string
	var l int
	for scanner.Scan() {
		oMap += scanner.Text() + "\n"
		l++
	}
	iMap := getBiggerMap(oMap, 80)
	/*

	   Right 1, down 1.
	   Right 3, down 1. (This is the slope you already checked.)
	   Right 5, down 1.
	   Right 7, down 1.
	   Right 1, down 2.

	*/
	fmt.Printf("oMap: %d, iMap: %d\n\n", l, len(strings.Split(iMap, "\n")))
	fmt.Printf("Right 1, down 1: %d\n", countTheTrees(iMap, 1, 1))
	fmt.Printf("Right 3, down 1 %d\n", countTheTrees(iMap, 1, 3))
	fmt.Printf("Right 5, down 1 %d\n", countTheTrees(iMap, 1, 5))
	fmt.Printf("Right 7, down 1 %d\n", countTheTrees(iMap, 1, 7))
	fmt.Printf("Right 1, down 2 %d\n", countTheTrees(iMap, 2, 1))
}

func getBiggerMap(oMap string, n int) string {
	var nm string
	for _, line := range strings.Split(oMap, "\n") {
		for i := 0; i < n; i++ {
			nm += line
		}
		nm += "\n"
	}
	nm = nm[:len(nm)-2]
	return nm
}

func countTheTrees(iMap string, down, right int) int {
	var x, y, c int
	lines := strings.Split(iMap, "\n")
	for y < len(lines) {
		if string(lines[y][x]) == "#" {
			c++
		}
		x += right
		y += down
	}
	return c
}

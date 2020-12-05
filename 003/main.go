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
	fmt.Printf("oMap: %d, iMap: %d\n\n", l, len(strings.Split(iMap, "\n")))
	fmt.Println(countTheTrees(iMap))
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

func countTheTrees(iMap string) int {
	var x, c int
	for _, line := range strings.Split(iMap, "\n") {
		if string(line[x]) == "#" {
			c++
		}
		x += 3
	}
	return c
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("second.txt")
	defer file.Close()

	p := make([]Password, 1000)
	scanner := bufio.NewScanner(file)
	var c int
	for scanner.Scan() {
		l := scanner.Text()
		parts := strings.Split(l, " ")
		ranges := strings.Split(parts[0], "-")
		min, _ := strconv.Atoi(ranges[0])
		max, _ := strconv.Atoi(ranges[1])
		p[c] = Password{
			Min:  min,
			Max:  max,
			Char: string(parts[1][0]),
			Pass: parts[2],
		}
		c++
	}
	fmt.Printf("Valid passwords: %d\n", CountValidPasswords(p))
	fmt.Printf("New Valid passwords: %d\n", CountNewValidPasswords(p))
}

type Password struct {
	Min  int
	Max  int
	Char string
	Pass string
}

func CountValidPasswords(passwords []Password) int {
	var c int
	for _, p := range passwords {
		if IsPassowrdValid(p) {
			c++
		}
	}
	return c
}

func CountNewValidPasswords(passwords []Password) int {
	var c int
	for _, p := range passwords {
		if IsPassValidNewRule(p) {
			c++
		}
	}
	return c
}

func IsPassowrdValid(p Password) bool {
	var sum int
	for _, c := range p.Pass {
		if string(c) == p.Char {
			sum++
		}
	}
	return sum >= p.Min && sum <= p.Max
}

func IsPassValidNewRule(p Password) bool {
	firstMatch := string(p.Pass[p.Min-1]) == p.Char
	secondMatch := string(p.Pass[p.Max-1]) == p.Char
	return firstMatch != secondMatch
}

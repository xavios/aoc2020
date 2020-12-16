package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("4.txt")
	p := ParsePassports(string(data))
	fmt.Print(CountValids(p))
}

type Passport struct {
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

func (p Passport) IsValid() bool {
	return p.BirthYear != 0 &&
		p.IssueYear != 0 &&
		p.ExpirationYear != 0 &&
		len(p.Height) != 0 &&
		len(p.HairColor) != 0 &&
		len(p.EyeColor) != 0 &&
		len(p.PassportID) != 0
}

func ParsePassports(data string) []Passport {
	lines := strings.Split(data, "\n")
	passports := make([]Passport, 0)
	var atLine int
	p := Passport{}
	for atLine < len(lines) {
		if len(strings.TrimSpace(lines[atLine])) == 0 {
			// next password will come or end of file
			passports = append(passports, p)
			p = Passport{}
		} else {
			p = parseLine(atLine, lines, p)
		}
		atLine++
	}
	passports = append(passports, p)
	return passports
}

func CountValids(p []Passport) int {
	var c int
	for _, pass := range p {
		if pass.IsValid() {
			c++
		}
	}
	return c
}

func parseLine(atLine int, lines []string, p Passport) Passport {
	entries := strings.Split(strings.TrimSpace(lines[atLine]), " ")
	for _, entry := range entries {
		kv := strings.Split(entry, ":")
		nv, _ := strconv.Atoi(kv[1])
		sv := kv[1]
		switch kv[0] {
		case "byr":
			p.BirthYear = nv
			break
		case "iyr":
			p.IssueYear = nv
			break
		case "eyr":
			p.ExpirationYear = nv
			break
		case "hgt":
			p.Height = sv
			break
		case "hcl":
			p.HairColor = sv
			break
		case "ecl":
			p.EyeColor = sv
			break
		case "pid":
			p.PassportID = sv
			break
		case "cid":
			p.CountryID = sv
			break
		}
	}
	return p
}

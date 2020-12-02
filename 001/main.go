package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"strconv"

	"github.com/pkg/errors"
)

func main() {
	account := make([]int, 0)

	file, err := os.Open("first.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("OMFG")
		}
		account = append(account, i)
	}

	mult, err := getMulAccounting(account, 2)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	mult2, _ := getMulAccounting(account, 3)
	fmt.Println(mult)
	fmt.Println(mult2)
}

func getMulAccounting(account []int, numOfParts int) (int, error) {
	if len(account) == 0 {
		return 0, errors.New("Empty account")
	}
	if numOfParts == 2 {
		for i := 0; i < len(account)-1; i++ {
			for j := i + 1; j < len(account); j++ {
				if account[i]+account[j] == 2020 {
					return account[i] * account[j], nil
				}
			}
		}
	} else if numOfParts == 3 {
		for i := 0; i < len(account)-2; i++ {
			for j := i + 1; j < len(account)-1; j++ {
				for k := j + 1; k < len(account); k++ {
					if account[i]+account[j]+account[k] == 2020 {
						return account[i] * account[j] * account[k], nil
					}
				}
			}
		}
	}
	return 0, nil
}

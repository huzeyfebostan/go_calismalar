package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func collatz(number int) int {
	if number%2 == 0 {
		number = number / 2
		return number
	} else {
		number = number*3 + 1
		return number
	}
	return 0
}

func main() {
	fmt.Println("Enter number: ")
	number := bufio.NewReader(os.Stdin)
	input, _ := number.ReadString('\n')
	input = strings.TrimSpace(input)
	num, _ := strconv.Atoi(input)

	fmt.Println("Enter number:")
	fmt.Println(num)
	collatznum := num
	for {
		collatznum = int(collatz(collatznum))
		fmt.Println(collatznum)
		if collatznum == 1 {
			break
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func findDigit(value string) int64 {

	var firstDigit rune
	var lastDigit rune

	runes := []rune(value)

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			firstDigit = runes[i]
			break
		}
	}

	for i := len(runes) - 1; i >= 0; i-- {
		if unicode.IsDigit(runes[i]) {
			lastDigit = runes[i]
			break
		}
	}

	num, _ := strconv.ParseInt(string(firstDigit)+string(lastDigit), 0, 0)

	fmt.Println(num)

	return num
}

func main() {
	f, err := os.OpenFile("input.txt", os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	filescanner := bufio.NewScanner(f)
	filescanner.Split(bufio.ScanLines)
	sum := 0
	for filescanner.Scan() {
		sum = sum + int(findDigit(filescanner.Text()))
	}

	fmt.Println(sum)
}

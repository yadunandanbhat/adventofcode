package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(path string) int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error() + " for path " + path)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), "")
		var index int = 0
		var num1 string
		var num2 string
		for idx, char := range arr {
			if _, err := strconv.ParseInt(char, 10, 32); err == nil {
				num1 = char
				index = idx
				break
			}
		}
		for idx := len(arr) - 1; idx > index; idx-- {
			if _, err := strconv.ParseInt(arr[idx], 10, 0); err == nil {
				num2 = arr[idx]
				index = idx
				break
			}
		}
		if len(num2) == 0 && len(num1) != 0 {
			num2 = num1
		}
		if len(num2) != 0 && len(num1) != 0 {
			if num, err := strconv.Atoi(num1 + num2); err == nil {
				sum += num
			}
		}
	}
	return sum
}

func main() {
	fmt.Println(readInput("C:\\Users\\Yadunandan Bhat\\Documents\\AOC2023\\01\\input.txt"))
}

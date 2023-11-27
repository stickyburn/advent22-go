package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("failed to open: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxCalories, currentElfCalories := 0, 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if currentElfCalories > maxCalories {
				maxCalories = currentElfCalories
			}

			currentElfCalories = 0
		} else {
			calories, _ := strconv.Atoi(line)
			currentElfCalories += calories
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("err reading: ", err)
	} else {
		fmt.Println("Max calories amoung elfs: ", maxCalories)
	}
}

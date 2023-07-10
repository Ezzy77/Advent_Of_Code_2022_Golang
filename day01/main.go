package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.Open("day01/sample.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	// Part 1
	currentCalories := 0
	maxCalories := 0

	for sc.Scan() {
		cal, err := strconv.Atoi(sc.Text())
		currentCalories += cal

		if err != nil {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
		}
	}
	fmt.Println(maxCalories)

}

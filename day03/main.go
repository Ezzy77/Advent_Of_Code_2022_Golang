package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	input, _ := os.Open("day03/sample.txt")
	defer input.Close()

	var rucksackList []string

	content := bufio.NewScanner(input)

	for content.Scan() {
		c := strings.Split(content.Text(), "\n")
		rucksackList = append(rucksackList, c...)
	}

	sum := findCommonItems(rucksackList)

	fmt.Println(sum)
}

func calculatePriority(itemType rune) int {
	if itemType >= 'a' && itemType <= 'z' {
		return int(itemType - 'a' + 1)
	} else if itemType >= 'A' && itemType <= 'Z' {
		return int(itemType - 'A' + 27)
	}
	return 0
}

func findCommonItems(rucksacks []string) int {
	totalPriority := 0

	for _, v := range rucksacks {
		leftCompartment := v[:len(v)/2]
		rightCompartment := v[len(v)/2:]

		commonItems := " "

		for _, l := range leftCompartment {

			if strings.ContainsRune(rightCompartment, l) {
				commonItems += string(l)

			}
		}

		for _, r := range commonItems {
			totalPriority += calculatePriority(r)

		}

	}
	return totalPriority
}

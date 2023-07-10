package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("day02/input.txt")
	defer input.Close()
	content := bufio.NewScanner(input)

	// Part 1
	myScore := 0
	rock := 1
	paper := 2
	scissors := 3
	winOutcome := 6
	drawOutcome := 3
	// part 2
	desiredScore := 0

	for content.Scan() {
		outcome := content.Text()

		switch outcome {
		case "A X":
			myScore = myScore + rock + drawOutcome
			desiredScore = desiredScore + scissors
		case "A Y":
			myScore = myScore + paper + winOutcome
			desiredScore = desiredScore + rock + drawOutcome

		case "A Z":
			myScore = myScore + scissors
			desiredScore = desiredScore + paper + winOutcome

		case "B X":
			myScore = myScore + rock
			desiredScore = desiredScore + rock

		case "B Y":
			myScore = myScore + paper + drawOutcome
			desiredScore = desiredScore + paper + drawOutcome

		case "B Z":
			myScore = myScore + scissors + winOutcome
			desiredScore = desiredScore + scissors + winOutcome

		case "C X":
			myScore = myScore + rock + winOutcome
			desiredScore = desiredScore + paper

		case "C Y":
			myScore = myScore + paper
			desiredScore = desiredScore + scissors + drawOutcome

		case "C Z":
			myScore = myScore + scissors + drawOutcome
			desiredScore = desiredScore + rock + winOutcome

		}

	}
	fmt.Println("Part 1 Results is:", myScore)
	fmt.Println("Part 2 Results is:", desiredScore)

}

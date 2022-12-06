package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var line string
	var points int
	scoreMap := map[string]int{
		"A X": 1 + 3,
		"A Y": 2 + 6,
		"A Z": 3 + 0,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 1 + 6,
		"C Y": 2 + 0,
		"C Z": 3 + 3,
	}
	// Part 2
	handmap := map[string]string{
		"A X": "A Z",
		"A Y": "A X",
		"A Z": "A Y",
		"B X": "B X",
		"B Y": "B Y",
		"B Z": "B Z",
		"C X": "C Y",
		"C Y": "C Z",
		"C Z": "C X",
	}

	file, err := os.OpenFile("input/day2.txt", os.O_RDONLY, 0755)
	if err != nil {
		fmt.Printf("Could not open file: %v\n", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Could not close file: %v\n", err)
		}
	}(file)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line = fileScanner.Text()
		// Solution for second part: convert input to same as before via handmap
		hand, _ := handmap[line]
		if score, ok := scoreMap[hand]; ok {
			points += score
		}
	}
	fmt.Printf("Total points: %d\n", points)
}

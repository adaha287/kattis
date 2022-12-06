package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var elfCalories []int
	var line string
	var cal int
	file, err := os.OpenFile("input/day1a.txt", os.O_RDONLY, 0755)
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
		if len(line) == 0 {
			elfCalories = append(elfCalories, cal)
			cal = 0
			fmt.Println("line was: empty")
		}
		if curCal, err := strconv.Atoi(line); err != nil {
		} else {
			cal += curCal
			fmt.Printf("line was: %v\n", line)
		}
	}
	sort.Ints(elfCalories)
	fmt.Printf("Elf with most calories: %d\n", elfCalories[len(elfCalories)-1])
	// Part b of the puzzle
	fmt.Printf("Total calories for the top 3 elfs with most calories: %d\n", elfCalories[len(elfCalories)-1]+elfCalories[len(elfCalories)-2]+elfCalories[len(elfCalories)-3])
}

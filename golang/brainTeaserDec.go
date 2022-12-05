package main

import (
	"fmt"
	"reflect"
	"sort"
)

func brainTeaserDec() {
	nums := []int{0, 2, 3, 5, 10, 15, 20, 25}
	sols := 0

	var uniqueSols [][]int
	solutions := make([][]int, 0)
	for _, v1 := range nums {
		for _, v2 := range nums {
			for _, v3 := range nums {
				if v1+v2+v3 == 30 {
					sols += 1
					fmt.Printf("%d: %d + %d + %d = 30\n", sols, v1, v2, v3)
					s := []int{v1, v2, v3}
					sort.Ints(s)
					solutions = append(solutions, s)
					if !itemExist(uniqueSols, s) {
						uniqueSols = append(uniqueSols, s)
					}
				}
			}
		}
	}
	fmt.Printf("Number of permutations: %d\n", sols)
	fmt.Printf("Number of unique solutions: %d\n", len(uniqueSols))
	for i, v := range uniqueSols {
		fmt.Printf("%d: %d, %d, %d\n", i+1, v[0], v[1], v[2])
	}

}

func itemExist(s [][]int, item []int) bool {
	for _, v := range s {
		if reflect.DeepEqual(v, item) {
			return true
		}
	}
	return false
}

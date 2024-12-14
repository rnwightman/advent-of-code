package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lists := make([][]int, 0)

	// parse input

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		inputValues := strings.Fields(line)

		if len(lists) == 0 {
			for range len(inputValues) {
				lists = append(lists, make([]int, 0))
			}
		}

		for index, inputValue := range inputValues {
			value, _ := strconv.Atoi(inputValue)
			lists[index] = append(lists[index], value)
		}
	}
	fmt.Println("input: ", lists)

	// sort lists

	for _, list := range lists {
		sort.Ints(list)
	}
	fmt.Println("sorted: ", lists)

	// calculate distances
	distances := make([]int, 0)
	totalDistance := 0

	for locNumber := range lists[0] {
		locations := make([]int, 0, len(lists))
		for i := range lists {
			locations = append(locations, lists[i][locNumber])
		}

		larger := slices.Max(locations)
		smaller := slices.Min(locations)
		distance := larger - smaller

		distances = append(distances, distance)
		totalDistance += distance
	}
	fmt.Println("distance:", distances)
	fmt.Println("total distance:", totalDistance)
}

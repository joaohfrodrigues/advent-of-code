package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

// Return the absolute value of x
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Part 1: get the total distance between the two lists
func getTotalDistance(list1 []int, list2 []int, inputLength int) int {
	// Sort the list1 and list2
	slices.Sort(list1)
	slices.Sort(list2)

	// Calculate the total distance
	totalDistance := 0
	for idx := range inputLength {
		totalDistance += abs(list1[idx] - list2[idx])
	}
	return totalDistance
}

// Count the number of elements in a slice
func count(slice []int, value int) int {
	count := 0
	for _, s := range slice {
		if s == value {
			count++
		}
	}
	return count
}

// Part 2: get the similarity between the two lists
// Similarity is defined as the product between the value of an element in the first list and the amount of times it appears in the second list
func getSimilarityScore(list1 []int, list2 []int, inputLength int) int {
	similarityScore := 0
	for idx := range inputLength {
		similarityScore += list1[idx] * count(list2, list1[idx])
	}
	return similarityScore
}

func main() {

	// Open and read input file
	file, err := os.Open("input.csv")

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
	}

	inputLength := len(records)

	// Initialize both lists
	var list1 []int
	var list2 []int
	list1 = make([]int, inputLength)
	list2 = make([]int, inputLength)

	// Populate both lists from records
	for idx, eachrecord := range records {
		list1[idx], _ = strconv.Atoi(eachrecord[0])
		list2[idx], _ = strconv.Atoi(eachrecord[1])
	}

	totalDistance := getTotalDistance(list1, list2, inputLength)
	fmt.Println("Part 1: total distance is", totalDistance)

	similarityScore := getSimilarityScore(list1, list2, inputLength)
	fmt.Println("Part 2: similarity score is", similarityScore)

}

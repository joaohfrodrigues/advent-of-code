package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func validateReport(report []int) bool {
	order := ""

	for idx, level := range report {
		if idx == 0 {
			if level < report[idx+1] {
				order = "asc"
			} else {
				order = "desc"
			}
		} else {
			if order == "asc" {
				if level <= report[idx-1] || level > report[idx-1] + 3 {
					return false
				}
			} else {
				if level >= report[idx-1] || level < report[idx-1] - 3 {
					return false
				}
			}
		}
	}
	return true
}

func removeIndex(s []int, index int) []int {
	returnSlice := make([]int, len(s))
	copy(returnSlice, s)
    return append(returnSlice[:index], returnSlice[index+1:]...)
}

func main() {
	// Open and read input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)


	var reports [][]int

    for fileScanner.Scan() {
        line := fileScanner.Text()
        parts := strings.Split(line, " ")
        intParts := make([]int, len(parts))
        for i, part := range parts {
            intPart, err := strconv.Atoi(part)
            if err != nil {
                log.Fatal(err)
            }
            intParts[i] = intPart
        }
        reports = append(reports, intParts)
    }

	safeReports := 0
	dampenedSafeReports := 0

    for _, report := range reports {
		isSafe := validateReport(report)
		
		if isSafe {
			safeReports++
		} else {
			for idx := range len(report) {
				filteredReport := removeIndex(report, idx)
				isDampenedSafe := validateReport(filteredReport)
				if isDampenedSafe {
					dampenedSafeReports++
					break
				}
			}
		}
    }

	fmt.Println("Part 1: Safe reports", safeReports)
	fmt.Println("Part 2: Dampened safe reports", safeReports + dampenedSafeReports)
}
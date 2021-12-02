package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func testSample() bool {
	test_input := []string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"}
	result := runSweep(test_input)

	if result != 5 {
		log.Fatalf("Expected 5 triplet increments, found %d", result)
		return false
	}

	return true

}

func loadInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func countIncrement(readings []int) int {
	incs := 0
	lastVal := readings[0]

	for _, current := range readings {
		if current > lastVal {
			incs++
		}
		lastVal = current
	}

	return incs
}

func groupInThree(readings []string) []int {
	sum := 0
	for i := 0; i < 2; i++ {
		current, err := strconv.Atoi(readings[i])
		if err != nil {
			panic(err)
		}
		sum += current
	}
	var groupSums []int
	groupSums = append(groupSums, sum)

	second, err := strconv.Atoi(readings[1])
	if err != nil {
		panic(err)
	}
	groupSums = append(groupSums, second)

	for i := 2; i < len(readings); i++ {
		current, err := strconv.Atoi(readings[i])
		if err != nil {
			panic(err)
		}

		groupSums[i-2] += current
		groupSums[i-1] += current
		groupSums = append(groupSums, current)
	}

	return groupSums
}

func runSweep(readings []string) int {
	grouped := groupInThree(readings)
	return countIncrement(grouped)
}

func main() {
	if testSample() {
		log.Println("Correct sample answer found!")
	}

	input, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	result := runSweep(input)
	log.Printf("Number of increments in input data is: %d", result)
}

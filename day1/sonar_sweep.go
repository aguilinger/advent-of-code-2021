package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func testSample() bool {
	test_input := []string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"}
	result := countDecrement(test_input)

	if result != 7 {
		log.Fatalf("Expected 7 decrements, found %d", result)
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

func countDecrement(readings []string) int {
	incs := 0
	lastVal, err := strconv.Atoi(readings[0])
	if err != nil {
		panic(err)
	}

	for _, r := range readings {
		current, err := strconv.Atoi(r)
		if err != nil {
			panic(err)
		}
		if current > lastVal {
			incs++
		}
		lastVal = current
	}

	return incs
}

func main() {
	if testSample() {
		log.Println("Correct sample answer found!")
	}

	input, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	result := countDecrement(input)
	log.Printf("Number of increments in input data is: %d", result)
}

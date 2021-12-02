package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func testSample() bool {
	test_input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	horiz, depth := position(test_input)
	if horiz != 15 && depth != 10 {
		log.Fatalf("Expected horizontal position 15 decrements, found %d; Expected depth 10, found %d", horiz, depth)
		return false
	}

	return true
}

func testSampleV2() bool {
	test_input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	horiz, depth := positionV2(test_input)
	if horiz != 15 && depth != 60 {
		log.Fatalf("Expected horizontal position 15 decrements, found %d; Expected depth 60, found %d", horiz, depth)
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

func position(course []string) (int, int) {
	s := regexp.MustCompile(` `)

	depth := 0
	horiz := 0
	for _, direction := range course {
		splitDir := s.Split(direction, 2)
		amount, err := strconv.Atoi(splitDir[1])
		if err != nil {
			panic(err)
		}

		switch splitDir[0] {
		case "forward":
			horiz += amount
		case "down":
			depth += amount
		case "up":
			depth -= amount
		}
	}

	return horiz, depth
}

func positionV2(course []string) (int, int) {
	s := regexp.MustCompile(` `)

	depth := 0
	horiz := 0
	aim := 0
	for _, direction := range course {
		splitDir := s.Split(direction, 2)
		amount, err := strconv.Atoi(splitDir[1])
		if err != nil {
			panic(err)
		}

		switch splitDir[0] {
		case "forward":
			horiz += amount
			depth += (aim * amount)
		case "down":
			aim += amount
		case "up":
			aim -= amount
		}
	}

	return horiz, depth
}

func main() {
	if testSample() && testSampleV2() {
		log.Println("Correct sample answer found!")
	}

	input, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}
	horiz, depth := position(input)
	log.Printf("Horizontal position: %d; Depth: %d; Product: %d", horiz, depth, horiz*depth)

	log.Println("Calculating with version 2...")
	horiz2, depth2 := positionV2(input)
	log.Printf("Horizontal position: %d; Depth: %d; Product: %d", horiz2, depth2, horiz2*depth2)
}

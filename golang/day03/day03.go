package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func output(ratings []int, batteries int) int {
	maxOutputs := make([]int, len(ratings))

	for _, rating := range ratings {
		for i := batteries; i >= 1; i-- {
			maxOutputs[i] = max(maxOutputs[i], maxOutputs[i-1]*10+rating)
		}
	}

	return maxOutputs[batteries]
}

func totalOutput(batteries int) int {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		ratings := make([]int, len(line))

		for i, char := range line {
			ratings[i] = int(char - '0')
		}

		sum += output(ratings, batteries)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return sum
}

func main() {
	fmt.Println("total output with 2 batteries:", totalOutput(2))
	fmt.Println("total output with 12 batteries:", totalOutput(12))
}

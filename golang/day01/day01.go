package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Dial struct {
	pos   int
	zeros int
}

func newDial() *Dial {
	return &Dial{pos: 50}
}

func (dial *Dial) rotate(dir string, dist int) {
	if dir == "L" {
		dial.pos -= dist
	} else {
		dial.pos += dist
	}

	dial.pos = mod(dial.pos, 100)

	if dial.pos == 0 {
		dial.zeros++
	}
}

func (dial *Dial) rotateClicks(dir string, dist int) {
	for i := 0; i < dist; i++ {
		dial.rotate(dir, 1)
	}
}

func rotateDial(clicks bool) int {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	dial := newDial()

	for scanner.Scan() {
		line := scanner.Text()

		dist, err := strconv.Atoi(line[1:])

		if err != nil {
			log.Fatalf("failed to parse integer: %s", err)
		}

		if clicks {
			dial.rotateClicks(line[:1], dist)
		} else {
			dial.rotate(line[:1], dist)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return dial.zeros
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func main() {
	fmt.Println("zeros:", rotateDial(false))
	fmt.Println("zeros with clicks:", rotateDial(true))
}

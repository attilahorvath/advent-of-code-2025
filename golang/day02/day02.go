package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func sumInvalidIds(line string, twiceOnly bool) int {
	ranges := strings.Split(line, ",")

	sum := 0

	for _, r := range ranges {
		interval := strings.Split(r, "-")

		min, err := strconv.Atoi(interval[0])

		if err != nil {
			log.Fatalf("failed to parse integer: %s", err)
		}

		max, err := strconv.Atoi(interval[1])

		if err != nil {
			log.Fatalf("failed to parse integer: %s", err)
		}

		for i := min; i <= max; i++ {
			id := strconv.Itoa(i)

			for l := 1; l <= len(id)/2; l++ {
				if len(id)%l != 0 {
					continue
				}

				if twiceOnly && (len(id)%2 != 0 || l != len(id)/2) {
					continue
				}

				invalid := true
				seq := ""

				for s := 0; s <= len(id)-l; s += l {
					sub := id[s:(s + l)]

					if seq != "" {
						if seq != sub {
							invalid = false
							break
						}
					} else {
						seq = sub
					}
				}

				if invalid {
					sum += i
					break
				}
			}
		}
	}

	return sum
}

func sumAllInvalidIds(twiceOnly bool) int {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	line := strings.TrimSuffix(string(data), "\n")

	return sumInvalidIds(line, twiceOnly)
}

func main() {
	fmt.Println("sum with sequences repeated twice only:", sumAllInvalidIds(true))
	fmt.Println("sum with all sequences:", sumAllInvalidIds(false))
}

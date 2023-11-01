package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	top := [3]int{0, 0, 0}
	var total int
	var count int

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			number, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			total = total + number
			continue
		}

		fmt.Printf("Elve %v: %v\n", count+1, total)
		if total > top[0] {
			top[2] = top[1]
			top[1] = top[0]
			top[0] = total
		} else if total > top[1] {
			top[2] = top[1]
			top[1] = total
		} else if total > top[2] {
			top[2] = total
		}
		total = 0
		count = count + 1
	}

	fmt.Printf("Elve %v: %v\n", count+1, total)
	if total > top[0] {
		top[2] = top[1]
		top[1] = top[0]
		top[0] = total
	} else if total > top[1] {
		top[2] = top[1]
		top[1] = total
	} else if total > top[2] {
		top[2] = total
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Top 3: %v\nMax calories: %v\n", top, top[0]+top[1]+top[2])
}

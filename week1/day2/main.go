package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	Day2_Pt1()
	Day2_Pt2()
}

func atoi(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}

func Day2_Pt1() {
	var validone, multiplier int
	sum := 0

	file, err := os.Open("day2/input.csv")

	if err != nil {
		log.Fatal("Error reading file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		log.Println("Error fouund here...")
		log.Fatal(err)
	}

	for _, record := range records {
		record_size := len(record)
		validone = 1

		if atoi(record[1])-atoi(record[0]) > 0 {
			multiplier = 1
		} else {
			multiplier = -1
		}

		for i := 1; i < record_size; i++ {
			diff := (atoi(record[i]) - atoi(record[i-1])) * multiplier
			if diff < 1 || diff > 3 {
				validone = 0
			}
		}

		sum += validone
	}

	fmt.Printf("total valid: %d\n", sum)

}

func Day2_Pt2() {
	var top, btm int
	sum := 0

	file, err := os.Open("day2/input.csv")

	if err != nil {
		log.Fatal("Error reading file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		log.Println("Error fouund here...")
		log.Fatal(err)
	}

	for _, record := range records {
		record_size := len(record)

		for i := 0; i < record_size; i++ {
			isValid := true
			inc_count := 0

			for j := 1; j < record_size; j++ {
				diff := 0

				top = j
				btm = j - 1

				if j == i {
					continue
				} else if j-1 == i {
					if j-1 == 0 {
						continue
					}
					btm = j - 2
				}

				diff = (atoi(record[top]) - atoi(record[btm]))
				absdiff := int(math.Abs(float64(diff)))

				if diff < 0 {
					inc_count--
				} else if diff > 0 {
					inc_count++
				}

				if absdiff < 1 || absdiff > 3 {
					isValid = false
					break
				}

			}

			inc_count = int(math.Abs(float64(inc_count)))

			if isValid && inc_count == (record_size-2) {
				sum++
				break
			}

		}

	}

	fmt.Printf("total valid: %d\n", sum)

}

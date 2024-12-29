package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	Day1_Pt1()
	Day1_Pt2()
}

func Day1_Pt1() {
	var arr1, arr2 [][]int
	index := 0
	sum := 0.0
	file, err := os.Open("day1/day1_input.csv")

	if err != nil {
		log.Fatal("Error reading file", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := regexp.MustCompile(`\s+`).Split(line, -1)
		field1, err1 := strconv.Atoi(fields[0])
		arr1 = append(arr1, []int{index, field1})
		field2, err2 := strconv.Atoi(fields[1])
		arr2 = append(arr2, []int{index, field2})

		if err1 != nil || err2 != nil {
			fmt.Println("Error appending...")
		}
		// fmt.Print(arr1[index])
		// fmt.Print(" ")
		// fmt.Println(arr2[index])
		index++
	}

	sort.Slice(arr1, func(a, b int) bool {
		if arr1[a][1] == arr1[b][1] {
			return arr1[a][0] < arr1[b][0]
		}
		return arr1[a][1] < arr1[b][1]
	})

	sort.Slice(arr2, func(a, b int) bool {
		if arr2[a][1] == arr2[b][1] {
			return arr2[a][0] < arr2[b][0]
		}
		return arr2[a][1] < arr2[b][1]
	})

	for i := 0; i < 1000; i++ {
		distance := math.Abs(float64(arr1[i][1]) - float64(arr2[i][1]))
		sum += distance
	}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Printf("%f\n", sum)

}

func Day1_Pt2() {
	var m1, m2 map[int]float64
	m1 = make(map[int]float64)
	m2 = make(map[int]float64)
	sum := 0.0
	file, err := os.Open("day1/day1_input.csv")

	if err != nil {
		log.Fatal("Error reading file", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := regexp.MustCompile(`\s+`).Split(line, -1)
		field1, err1 := strconv.Atoi(fields[0])
		m1[field1] += 1
		field2, err2 := strconv.Atoi(fields[1])
		m2[field2] += 1

		if err1 != nil || err2 != nil {
			fmt.Println("Error appending...")
		}

	}

	for key, _ := range m1 {
		if m2[key] > 0 {
			fmt.Printf("%.0f/%.0f, ", m2[key], float64(key))
			sum += float64(key) * m2[key]
		}
	}
	fmt.Println()

	fmt.Printf("%f\n", sum)

}

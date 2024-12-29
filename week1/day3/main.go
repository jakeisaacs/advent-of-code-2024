package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	// "math"
	"os"
	"strconv"
)

func main() {
	Day3_Pt1()
	Day3_Pt2()
}

func atoi(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}

func ReadFile(fname string) string {
	file, err := os.ReadFile(fname)

	if err != nil {
		log.Fatal("Error reading file", err)
	}

	// defer file.Close()

	return string(file)
}

func Day3_Pt1() {
	var n1, n2 int
	file := ReadFile("day3/input.txt")
	sum := 0

	r2 := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := r2.FindAllStringSubmatch(file, -1)

	for _, m := range matches {
		n1 = strings.Index(m[0], ",")
		n2 = strings.Index(m[0], ")")
		sum += atoi(m[0][4:n1]) * atoi(m[0][n1+1:n2])
		// fmt.Printf("%s - %s - %s - %d\n", m[0], m[0][4:n1], m[0][n1+1:n2], sum)
	}

	fmt.Printf("Total: %d\n", sum)
}

func Day3_Pt2() {
	var n1, n2 int
	file := ReadFile("day3/input.txt")
	sum := 0

	r1 := regexp.MustCompile(`don't\(\)(.|\n)*?do\(\)`)

	file = r1.ReplaceAllString(file, "")
	// fmt.Println(file)

	r2 := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := r2.FindAllStringSubmatch(file, -1)

	for _, m := range matches {
		n1 = strings.Index(m[0], ",")
		n2 = strings.Index(m[0], ")")
		sum += atoi(m[0][4:n1]) * atoi(m[0][n1+1:n2])
		// fmt.Printf("%s - %s - %s - %d\n", m[0], m[0][4:n1], m[0][n1+1:n2], sum)
	}

	fmt.Printf("Total: %d\n", sum)
}

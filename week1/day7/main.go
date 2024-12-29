package main

import (
	"bufio"
	"fmt"
	"log"

	// "math"
	"os"
	"strconv"
)

func main() {
	Part1()
	Part2()
}

func atoi(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}

func ReadFile(fname string) []string {
	var grid []string
	file, err := os.Open(fname)

	if err != nil {
		log.Fatal("Error reading file", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
		// fmt.Println(line)
	}

	return grid
}

func CheckRight(grid *[]string, letter string, i int, j int, clim int) int {
	if j >= clim {
		return 0
	}

	word := letter + string((*grid)[i][j]) + string((*grid)[i][j+1]) + string((*grid)[i][j+2])

	if word == "XMAS" || word == "SAMX" {
		return 1
	}

	return 0
}

func CheckDown(grid *[]string, letter string, i int, j int, rlim int) int {
	if i >= rlim {
		return 0
	}

	word := letter + string((*grid)[i][j]) + string((*grid)[i+1][j]) + string((*grid)[i+2][j])

	if word == "XMAS" || word == "SAMX" {
		return 1
	}

	return 0
}

func CheckDiagLeft(grid *[]string, letter string, i int, j int, rlim int) int {
	if j < 2 || i >= rlim {
		return 0
	}

	word := letter + string((*grid)[i][j]) + string((*grid)[i+1][j-1]) + string((*grid)[i+2][j-2])

	if word == "XMAS" || word == "SAMX" {
		return 1
	}

	return 0
}

func CheckDiagRight(grid *[]string, letter string, i int, j int, clim int, rlim int) int {
	if j >= clim || i >= rlim {
		return 0
	}

	word := letter + string((*grid)[i][j]) + string((*grid)[i+1][j+1]) + string((*grid)[i+2][j+2])

	if word == "XMAS" || word == "SAMX" {
		return 1
	}

	return 0
}

func Part2() {
	// var p1, p2 *string
	grid := ReadFile("day4/input.txt")
	rows := len(grid)
	cols := len(grid[0])
	sum := 0

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if grid[i][j] == 'A' {
				d1 := string(grid[i-1][j-1]) + string(grid[i+1][j+1])
				d2 := string(grid[i+1][j-1]) + string(grid[i-1][j+1])
				if (d1 == "MS" || d1 == "SM") && (d2 == "MS" || d2 == "SM") {
					sum++
				}
			}
		}
	}

	fmt.Printf("Total: %d\n", sum)

}

func Part1() {
	// var p1, p2 *string
	grid := ReadFile("day4/input.txt")
	rows := len(grid)
	cols := len(grid[0])
	rlim := rows - 2
	clim := cols - 2
	sum := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'X' || grid[i][j] == 'S' {
				sum += CheckRight(&grid, string(grid[i][j]), i, j+1, clim)
				sum += CheckDown(&grid, string(grid[i][j]), i+1, j, rlim)
				sum += CheckDiagLeft(&grid, string(grid[i][j]), i+1, j-1, rlim)
				sum += CheckDiagRight(&grid, string(grid[i][j]), i+1, j+1, clim, rlim)
			}
		}
	}

	fmt.Printf("Total: %d\n", sum)

}

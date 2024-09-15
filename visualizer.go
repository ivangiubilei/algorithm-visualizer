package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"time"
)

func move(arr *[]int, i, j int) {
	if len(*arr) > i && len(*arr) > j && i != j {
		tmp := (*arr)[i]
		(*arr)[i] = (*arr)[j]
		(*arr)[j] = tmp
	}
}

func mix(arr *[]int, times int) {
	for n := 0; n < times; n++ {
		for i := range *arr {
			move(arr, i, rand.IntN(len(*arr)))
		}
	}
}

func printArray(arr []int, comparisons, swaps int, status bool, name string) {
	clearScreen()
	for i := 0; i < len(arr); i++ {
		for j := 0; j < arr[i]; j++ {
			if j == 0 {
				fmt.Print(" |")
			}
			fmt.Print("█")
		}
		fmt.Println()
	}
	if status {
		fmt.Printf("\n\n%s\n", name)
		fmt.Printf("----------------\n")
		fmt.Printf("Comparisons: %d\n", comparisons)
		fmt.Printf("Swaps: %d\n", swaps)
	}
}

func printArrayCurrent(arr []int, current, line, comparisons, swaps int, status bool, name string) {
	var Reset = "\033[0m"
	var Red = "\033[31m"
	var Blue = "\033[34m"

	clearScreen()
	for i := 0; i < len(arr); i++ {
		for j := 0; j < arr[i]; j++ {
			if j == 0 {
				fmt.Print(" |")
			}
			if i == current {
				fmt.Print(Red + "▒" + Reset)
			} else if i == line {
				fmt.Print(Blue + "▒" + Reset)
			} else {
				fmt.Print("█")
			}
		}
		fmt.Println()
	}
	if status {
		fmt.Printf("\n\n%s\n", name)
		fmt.Printf("----------------\n")
		fmt.Printf("Comparisons: %d\n", comparisons)
		fmt.Printf("Swaps: %d\n", swaps)
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func BubbleSort(arr *[]int, ms int, status bool) {
	comparisons := 0
	swaps := 0
	algName := "BubbleSort"
	printArray(*arr, 0, 0, false, algName)
	for i := range *arr {
		time.Sleep(time.Duration(ms) * time.Millisecond)
		for j := range *arr {
			comparisons++
			time.Sleep(time.Duration(ms) * time.Millisecond)
			if (*arr)[i] > (*arr)[j] {
				swaps++
				move(arr, i, j)
			}
			printArrayCurrent(*arr, i, j, comparisons, swaps, status, algName)
		}
	}
	printArray(*arr, comparisons, swaps, status, algName)
}

func main() {
	size := flag.Int("size", 10, "An array of the specified size will be created")
	ms := flag.Int("ms", 50, "This will be the sleep time used for each iteration")
	status := flag.Bool("status", true, "Show algorithm name, number of comparison and number of swaps")

	flag.Parse()
	var arrayToSort []int = make([]int, *size)

	for indx := range arrayToSort {
		arrayToSort[indx] = indx + 1
	}

	mix(&arrayToSort, 10)

	BubbleSort(&arrayToSort, *ms, *status)
}

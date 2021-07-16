package main

import (
	"fmt"
	"sort"
)

// function that returns an average value of array
func arrays(s [6]int) float64 {
	var sum float64 = 0
	for _, number := range s {
		x := float64(number)
		sum += x
	}
	n := sum / float64(len(s))
	return n
}

// function, that returns the longest word from the slice of strings
func maxString(words []string) string {
	var maxWord string
	for _, word := range words {
		if len(word) > len(maxWord) {
			maxWord = word
		}
	}
	return maxWord
}

// function, that returns the longest word from the slice of strings
func maxStringBySort(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	s := words[len(words)-1]
	return s
}

// function, that returns the copy of the original slice in reverse order
func reverse(numbers []int64) []int64 {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

// function, that prints map values sorted in order of increasing keys
func printSorted(m map[int]string) {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	var keys1 []string
	for _, k := range keys {
		keys1 = append(keys1, m[k])
	}
	fmt.Println("Task_4:", keys1)
}

func main() {
	z := [6]int{1, 2, 3, 4, 5, 6}
	y := arrays(z)
	fmt.Println("Task_1:", y)

	words1 := []string{"one", "two", "three"}
	w := maxString(words1)
	fmt.Println("Task_2.1:", w)

	w = maxStringBySort(words1)
	fmt.Println("Task_2.2:", w)

	words2 := []string{"one", "two"}
	maxWord1 := maxString(words2)
	fmt.Println("Task_2.3:", maxWord1)

	r := []int64{1, 2, 5, 15}
	n := reverse(r)
	fmt.Println("Task_3:", n)

	key1 := map[int]string{2: "a", 0: "b", 1: "c"}
	printSorted(key1)

	key2 := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	printSorted(key2)
}

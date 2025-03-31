package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"
	"sort"
	"strings"
)

// What are the N most common words in sherlock.txt

// `a` is a "raw" string, at \ is just a \
var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

// Code that runs before main
// - var expressions
// - init function

func main() {
	// mapDemo()
	file, err := os.Open("sherlock.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer file.Close()

	freq := make(map[string]int) // word -> count
	s := bufio.NewScanner(file)
	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1)
		for _, word := range words {
			freq[strings.ToLower(word)]++
		}
	}
	if err := s.Err(); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	top := topN(freq, 10)
	fmt.Println(top)
}

// topN returns the "n" most common words in freq.
func topN(freq map[string]int, n int) []string {
	words := slices.Collect(maps.Keys(freq))
	sort.Slice(words, func(i, j int) bool {
		wi, wj := words[i], words[j]
		// Sort in reverse order.
		return freq[wi] > freq[wj]
	})

	n = min(n, len(words))
	return words[:n]
}

func mapDemo() {
	heros := map[string]string{ // hero -> name
		"Superman":     "Clark",
		"Wonder Woman": "Diana",
		"Batman":       "Bruce",
	}

	// Keys
	for k := range heros {
		fmt.Println(k)
	}

	// Key + Value
	for k, v := range heros {
		fmt.Println(v, "is", k)
	}

	// Values
	for _, v := range heros {
		fmt.Println(v)
	}

	n := heros["Batman"]
	fmt.Println(n)

	// Accessing non-existing key return zero value
	n = heros["Aquaman"]
	fmt.Printf("%q\n", n)

	// Use "comma ok" to find if a key is in the map
	n, ok := heros["Aquaman"]
	if ok {
		fmt.Printf("%q\n", n)
	} else {
		fmt.Println("Aquaman not found")
	}

	delete(heros, "Batman")
	fmt.Println(heros)
}

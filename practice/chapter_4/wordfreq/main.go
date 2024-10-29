package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	input, err := os.Open("words.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v", err)
		os.Exit(1)
	}
	defer input.Close()

	counts := make(map[string]int)

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	re := regexp.MustCompile(`[^\w\s]`)

	for scanner.Scan() {
		word := scanner.Text()

		word = strings.ToLower(word)

		word = re.ReplaceAllString(word, "")

		if word == "" {
			continue
		}

		counts[word]++
	}

	for k, v := range counts {
		if v != 1 {
			fmt.Printf("Word %q appears %d times\n", k, v)
		} else {
			fmt.Printf("Word %q appears 1 time\n", k)
		}
	}

	fmt.Printf("\nThe word 'rain' appears %v times\n", counts["rain"])

}

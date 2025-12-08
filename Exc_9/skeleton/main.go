package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"exc9/mapred"
)

// Main function
func main() {
	// todo read file
	file, err := os.Open("res/meditations.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// todo run your mapreduce algorithm
	var mr mapred.MapReduce
	results := mr.Run(lines)
	// todo print your result to stdout
	for k, v := range results {
		fmt.Printf("%s: %d\n", k, v)
	}
}

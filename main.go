package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	log.SetFlags(0)

	pattern := flag.String("pattern", "", "Regex pattern to test")
	flag.Parse()

	if *pattern == "" {
		log.Fatalf("Usage: regex-util --pattern='<regex>'")
	}

	re, err := regexp.Compile(*pattern)
	if err != nil {
		log.Fatalf("Invalid regex: %v\n", err)
	}

	log.Printf("Compiled regex: %s\n", *pattern)
	log.Println("Enter strings to test:")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		matchIndices := re.FindAllStringIndex(input, -1)
		if len(matchIndices) == 0 {
			log.Println("No match")
		} else {
			highlighted := highlightMatches(input, matchIndices)
			log.Printf("Match: %s\n", highlighted)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading input: %v\n", err)
	}
}

func highlightMatches(s string, matchIndices [][]int) string {
	var sb strings.Builder
	last := 0

	for _, pair := range matchIndices {
		start, end := pair[0], pair[1]
		sb.WriteString(s[last:start])
		sb.WriteString("\x1b[32m") // Highlight match as green
		sb.WriteString(s[start:end])
		sb.WriteString("\x1b[0m") // Reset colour
		last = end
	}
	sb.WriteString(s[last:])

	return sb.String()
}

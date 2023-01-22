package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"
)

func process(flags Flags) {
	inputFile := flags.inputFile
	outputFile := flags.outputFile
	regex := flags.regex
	threads := flags.threads

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	re := regexp.MustCompile(regex)

	rePatterns := re.SubexpNames()[1:]
	patternsLength := len(rePatterns)
	rePatterns = append(rePatterns, "garbage")

	writer, err := ParallelCsvWriter(outputFile)
	if err != nil {
		panic(err)
	}

	writer.Write(rePatterns)

	fmt.Printf("CSV header: %s\n", rePatterns)
	fmt.Printf("Regex: %s\n", re)
	fmt.Printf("Threads: %d\n", threads)

	var wg sync.WaitGroup

	lines := make(chan string)

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func(i int) {
			var bufferedLines [][]string

			for line := range lines {
				match := re.FindStringSubmatch(line)

				if len(match) == (patternsLength + 1) {
					bufferedLines = append(bufferedLines, match[1:])
				} else {
					bufferedLines = append(bufferedLines, append(make([]string, patternsLength), line))
				}

				if len(bufferedLines)%15_000 == 0 {
					writer.WriteAll(bufferedLines)
					bufferedLines = [][]string{}
				}
			}

			writer.WriteAll(bufferedLines)
			wg.Done()
		}(i)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines <- scanner.Text()
	}

	close(lines)
	wg.Wait()
}

func main() {
	flags := GetFlags()
	CheckFlags(flags)

	process(flags)

	fmt.Println("Done!")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"
)

var numWorkers = 48

func process(inputFile string, outputFile string, regex string) {
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
	rePatterns = append(rePatterns, "garbage")

	writer, err := ParallelCsvWriter(outputFile)
	if err != nil {
		panic(err)
	}

	writer.Write(rePatterns)

	fmt.Printf("CSV header: %s\n", rePatterns)
	fmt.Printf("Regex: %s\n", re)

	var wg sync.WaitGroup

	lines := make(chan string)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(i int) {
			var bufferedLines [][]string

			for line := range lines {
				match := re.FindStringSubmatch(line)

				if len(match) == 5 {
					bufferedLines = append(bufferedLines, match[1:])
				} else {
					bufferedLines = append(bufferedLines, []string{"", "", "", "", line})
				}

				if len(bufferedLines)%15_000 == 0 {
					writer.WriteAll(bufferedLines)
					//fmt.Printf("Thread #%02d >> added %d lines\n", i, len(bufferedLines))

					bufferedLines = [][]string{}
				}
			}

			writer.WriteAll(bufferedLines)
			//fmt.Printf("Thread #%02d >> added %d lines\n", i, len(bufferedLines))

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
	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s <input_file> <output_file> <regex>\n", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	process(os.Args[1], os.Args[2], os.Args[3])

	fmt.Println("Done!")
}

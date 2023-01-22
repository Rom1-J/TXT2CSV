package main

import "flag"

type Flags struct {
	inputFile  string
	outputFile string
	regex      string
	threads    int
}

func GetFlags() Flags {
	inputFile := flag.String("input", "", "Input file")
	outputFile := flag.String("output", "output.csv", "Output file")
	regex := flag.String("regex", "", "Regex to use")

	threads := flag.Int("threads", 12, "Number of threads to use")

	flag.Parse()

	return Flags{
		inputFile:  *inputFile,
		outputFile: *outputFile,
		regex:      *regex,
		threads:    *threads,
	}
}

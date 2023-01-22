package utils

import "flag"

type Flags struct {
	InputFile  string
	OutputFile string
	Regex      string
	Threads    int
}

func GetFlags() Flags {
	inputFile := flag.String("input", "", "Input file")
	outputFile := flag.String("output", "output.csv", "Output file")
	regex := flag.String("regex", "", "Regex to use")

	threads := flag.Int("threads", 12, "Number of threads to use")

	flag.Parse()

	return Flags{
		InputFile:  *inputFile,
		OutputFile: *outputFile,
		Regex:      *regex,
		Threads:    *threads,
	}
}

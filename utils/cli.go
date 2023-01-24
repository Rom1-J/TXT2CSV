package utils

import (
	"flag"
	"runtime"
	"strings"
)

type Flags struct {
	InputFile  string
	OutputFile string
	Regex      string
	Threads    int
}

func GetFlags() Flags {
	inputFile := flag.String("input", "", "Input file")
	outputFile := flag.String("output", "", "Output file (default \"stdout\")")
	regex := flag.String("regex", "", "Regex to use")

	threads := flag.Int("threads", 12, "Number of threads to use")

	flag.Parse()

	if runtime.GOOS == "windows" {
		return Flags{
			InputFile:  strings.TrimRight(*inputFile, `'"`),
			OutputFile: *outputFile,
			Regex:      *regex,
			Threads:    *threads,
		}
	}

	return Flags{
		InputFile:  *inputFile,
		OutputFile: *outputFile,
		Regex:      *regex,
		Threads:    *threads,
	}
}

package utils

import (
	"fmt"
	"os"
	"regexp"
)

func CheckFlags(flags Flags) {
	// Check input file exists
	file, err := os.Open(flags.InputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// Check regex has named patterns
	re := regexp.MustCompile(flags.Regex)
	if len(re.SubexpNames()[1:]) == 0 {
		fmt.Println("Regex does not contain groups or named patterns.")
		os.Exit(1)
	}

	// Check threads positivity
	if flags.Threads <= 0 {
		fmt.Println("Threads must be 1 or greater.")
		os.Exit(1)
	}
}

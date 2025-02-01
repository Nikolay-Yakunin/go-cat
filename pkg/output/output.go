package output

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/Nikolay-Yakunin/go-cat/pkg/flags"
	"os"
)

// ProcessFile reads the given file line by line and processes each line based on the provided flags.
func ProcessFile(file *os.File, flags flags.Flags) {
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	previousBlank := false

	for scanner.Scan() {
		line := scanner.Text()
		printLine(line, flags, &lineNumber, &previousBlank)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %q: %v\n", file.Name(), err)
	}
}

// printLine prints a single line according to the flags.
func printLine(rawLine string, flags flags.Flags, lineNumber *int, previousBlank *bool) {
	if flags.FlagS && *previousBlank && len(rawLine) == 0 {
		return
	}

	shouldNumber := false
	if flags.FlagB {
		if len(rawLine) > 0 {
			shouldNumber = true
		}
	} else if flags.FlagN {
		shouldNumber = true
	}

	line := rawLine

	if flags.FlagT {
		line = strings.ReplaceAll(line, "\t", "^I")
	}

	if flags.FlagE {
		line += "$"
	}

	if shouldNumber {
		fmt.Printf("%6d\t%s\n", *lineNumber, line)
		(*lineNumber)++
	} else {
		fmt.Println(line)
	}

	*previousBlank = (len(rawLine) == 0)
}

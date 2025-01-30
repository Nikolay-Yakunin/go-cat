package output

import (
	"bufio"
	"fmt"
	"github.com/Nikolay-Yakunin/go-cat/pkg/flags"
	"os"
)

func ProcessFile(file *os.File, flags flags.Flags) {
	scanner := bufio.NewScanner(file)
	lineCount := 1
	var isPreviousState bool

	for scanner.Scan() {
		line := scanner.Text()
		PrintLine(line, flags, &lineCount, &isPreviousState)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
	}
}

func PrintLine(line string, flags flags.Flags, lineCount *int, isPreviousState *bool) {
	if flags.FlagS && *isPreviousState && len(line) == 0 {
		return
	}

	if flags.FlagB && len(line) > 0 && (len(line) == 0 || line[0] != '\n') {
		fmt.Printf("%6d\t%s\n", *lineCount, line)
		*lineCount++
	} else if flags.FlagN && !flags.FlagB {
		fmt.Printf("%6d\t%s\n", *lineCount, line)
		*lineCount++
	} else {
		fmt.Println(line)
	}

	*isPreviousState = (len(line) == 0)
}

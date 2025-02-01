package main

import (
	"fmt"
	"os"

	"github.com/Nikolay-Yakunin/go-cat/pkg/flags"
	"github.com/Nikolay-Yakunin/go-cat/pkg/output"
)

func main() {
	// Initialize the flag state with default values.
	flagState := &flags.Flags{
		FlagV: false,
		FlagB: false,
		FlagE: false,
		FlagN: false,
		FlagS: false,
		FlagT: false,
		FlagH: false,
	}

	// Parse command-line arguments into flags and a slice of file names.
	fileNames := flags.ParseArgs(os.Args, flagState)

	// add help flag
	// use "os.Args[0]" because user can change name
	if flagState.FlagH {
		printHelp(os.Args[0])
		os.Exit(0)
	}

	// If no file is provided, print usage and exit.
	if len(fileNames) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... [FILE]...\n", os.Args[0])
		os.Exit(1)
	}

	// Process each file sequentially.
	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %q: %v\n", fileName, err)
			continue
		}

		// Process the file with the provided output package.
		output.ProcessFile(file, *flagState)

		// Close the file immediately after processing.
		if err = file.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Error closing file %q: %v\n", fileName, err)
		}
	}
}


func printHelp(progName string) {
	helpMessage := `Usage: ` + progName + ` [OPTION]... [FILE]...

A simple implementation of the cat command with additional options.

Options:
  -b, --number-nonblank     Number nonempty output lines, overrides -n.
  -e, -E                   Display a '$' at end of each line.
  -n, --number             Number all output lines.
  -s, --squeeze-blank      Suppress repeated empty output lines.
  -t, -T                   Display tab characters as '^I'.
  -v                      Display non-printing characters (with -e and -t effects).
  -h, --help               Display this help and exit.

Examples:
  ` + progName + ` -n file.txt
  ` + progName + ` --number-nonblank file1.txt file2.txt
`
	fmt.Println(helpMessage)
}

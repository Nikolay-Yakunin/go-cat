package flags

import (
	"fmt"
	"os"
	"strings"
)

// Flags holds the state of all command-line flags.
type Flags struct {
	FlagV bool // Display non-printing characters visibly.
	FlagB bool // Number non-empty output lines.
	FlagE bool // Display a '$' at the end of each line.
	FlagN bool // Number all output lines.
	FlagS bool // Suppress repeated empty output lines.
	FlagT bool // Display tab characters as '^I'.
	FlagH bool // Display help message.
}

func ParseLongFlag(arg string, flags *Flags) bool {
	switch arg {
	case "--number-nonblank":
		flags.FlagB = true
	case "--squeeze-blank":
		flags.FlagS = true
	case "--number":
		flags.FlagN = true
	case "--help":
		flags.FlagH = true
	default:
		return false
	}
	return true
}

// ParseShortFlags processes a combined set of short flags (e.g. -ben).
// Unknown flags are reported to standard error.
func ParseShortFlags(arg string, flags *Flags) {
	// Iterate over each character after the leading '-' character.
	for _, char := range arg[1:] {
		switch char {
		case 'b':
			flags.FlagB = true
		case 'E', 'e':
			flags.FlagE = true
			flags.FlagV = true
		case 'n':
			flags.FlagN = true
		case 's':
			flags.FlagS = true
		case 'T', 't':
			flags.FlagT = true
			flags.FlagV = true
		case 'v':
			flags.FlagV = true
		case 'h':
			flags.FlagH = true
		default:
			PrintUnknownOption(char)
		}
	}
}

// PrintUnknownOption prints a message for an unrecognized flag.
func PrintUnknownOption(unknownFlag rune) {
	fmt.Fprintf(os.Stderr, "Unknown option: -%c\n", unknownFlag)
}

// ParseArgs processes command-line arguments and separates flags from file names.
// It follows the convention that if an argument is "--", all subsequent arguments are file names,
// and any argument that does not start with '-' (or is exactly "-") is treated as a file name.
func ParseArgs(argv []string, flags *Flags) []string {
	var files []string
	parsingFlags := true

	// Iterate over all arguments, skipping the program name at index 0.
	for _, arg := range argv[1:] {
		if parsingFlags {
			// If argument is "--", disable flag parsing for the rest.
			if arg == "--" {
				parsingFlags = false
				continue
			}
			// If argument starts with '-' and is not exactly "-", treat it as flag(s).
			if strings.HasPrefix(arg, "-") && arg != "-" {
				if strings.HasPrefix(arg, "--") {
					if !ParseLongFlag(arg, flags) {
						// Report unknown long option by printing the first character after "--".
						PrintUnknownOption(rune(arg[2]))
					}
				} else {
					ParseShortFlags(arg, flags)
				}
				// Continue parsing flags.
				continue
			}
		}
		// If flag parsing is disabled or the argument doesn't look like a flag, treat it as a file name.
		files = append(files, arg)
	}
	return files
}

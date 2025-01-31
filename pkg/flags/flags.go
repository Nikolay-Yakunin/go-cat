package flags

import (
	"fmt"
	"os"
	"strings"
)

type Flags struct {
	FlagV bool
	FlagB bool
	FlagE bool
	FlagN bool
	FlagS bool
	FlagT bool
	FlagH bool
}

type Args struct {
	First int
	Last  int
}

func ParseLongFlag(arg string, flags *Flags) bool {
	if arg == "--number-nonblank" {
		flags.FlagB = true
	} else if arg == "--squeeze-blank" {
		flags.FlagS = true
	} else if arg == "--number" {
		flags.FlagN = true
	} else {
		return false
	}
	return true
}

func ParseShortFlags(arg string, flags *Flags) {
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
		default:
			PrintUnknownOption(char)
		}
	}
}

func PrintUnknownOption(unknownFlag rune) {
	fmt.Fprintf(os.Stderr, "Unknown option: -%c\n", unknownFlag)
}

func ParseFlags(argc int, argv []string, flags *Flags) Args {
	var file, last, f int = -1, -1, 1
	args := Args{First: -1, Last: -1}

	for i := 1; i < argc; i++ {
		if strings.HasPrefix(argv[i], "-") {
			if strings.HasPrefix(argv[i], "--") {
				if !ParseLongFlag(argv[i], flags) {
					PrintUnknownOption(rune(argv[i][0]))
				}
			} else {
				ParseShortFlags(argv[i], flags)
			}
		} else {
			if f == 1 {
				f = 0
				file = i
			}
			last = i
		}
	}

	if file == -1 {
		file = argc
	}
	// Fix this shit
	args.First = file
	args.Last = last

	return args
}

package main

import (
	"fmt"
	"os"

	"github.com/Nikolay-Yakunin/go-cat/pkg/flags"
	"github.com/Nikolay-Yakunin/go-cat/pkg/output"
)

func main() {
	argc := len(os.Args)
	argv := os.Args

	flags_state := &flags.Flags{
		FlagV: false,
		FlagB: false,
		FlagE: false,
		FlagN: false,
		FlagS: false,
		FlagT: false,
		FlagH: false,
	}

	args := flags.ParseFlags(argc, argv, flags_state)

	if args.First == argc {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... [FILE]...\n", argv[0])
		os.Exit(1)
	}

	for i := args.First; i <= args.Last; i++ {
		file, err := os.Open(argv[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %s\n", err)
			continue
		}
		defer file.Close()
		output.ProcessFile(file, *flags_state)
	}
}

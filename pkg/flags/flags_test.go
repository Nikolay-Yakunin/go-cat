package flags

import (
	"testing"
)

func TestParseLongFlag(t *testing.T) {
	flags := &Flags{}
	if !ParseLongFlag("--number-nonblank", flags) || !flags.FlagB {
		t.Error("Expected flag_b to be true")
	}
	if !ParseLongFlag("--squeeze-blank", flags) || !flags.FlagS {
		t.Error("Expected flag_s to be true")
	}
	if !ParseLongFlag("--number", flags) || !flags.FlagN {
		t.Error("Expected flag_n to be true")
	}
	if ParseLongFlag("--invalid", flags) {
		t.Error("Expected invalid flag to return false")
	}
}

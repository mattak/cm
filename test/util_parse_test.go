package main

import (
	"github.com/mattak/cm/cmd"
	"testing"
)

func TestTrimEmptyLines(t *testing.T) {
	newLines1 := cmd.TrimEmptyLines([]string{"{}", "#", "# ", ""})
	if len(newLines1) != 1 {
		t.Fatalf("expect1 failed: %d", len(newLines1))
	}

	newLines2 := cmd.TrimEmptyLines([]string{})
	if len(newLines2) != 0 {
		t.Fatal("expect2 failed")
	}
}

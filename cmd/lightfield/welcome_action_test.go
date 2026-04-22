package main

import (
	"os"
	"testing"
)

func TestShouldShowWelcomeHonorsEnvironmentOptOuts(t *testing.T) {
	stdin, err := os.Open(os.DevNull)
	if err != nil {
		t.Fatalf("open stdin stub: %v", err)
	}
	defer stdin.Close()

	stdout, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	if err != nil {
		t.Fatalf("open stdout stub: %v", err)
	}
	defer stdout.Close()

	t.Setenv("CI", "true")
	if shouldShowWelcome(stdin, stdout) {
		t.Fatal("expected welcome to be disabled in CI")
	}

	t.Setenv("CI", "")
	t.Setenv("LIGHTFIELD_NO_WELCOME", "1")
	if shouldShowWelcome(stdin, stdout) {
		t.Fatal("expected welcome to be disabled when opted out")
	}

	t.Setenv("LIGHTFIELD_NO_WELCOME", "")
	t.Setenv("TERM", "dumb")
	if shouldShowWelcome(stdin, stdout) {
		t.Fatal("expected welcome to be disabled for dumb terminals")
	}
}

package main

import (
	"os"
	"testing"
)

func TestShouldShowWelcome(t *testing.T) {
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

	for _, tc := range []struct {
		name string
		env  map[string]string
	}{
		{"CI set", map[string]string{"CI": "true"}},
		{"opt-out env var", map[string]string{"LIGHTFIELD_NO_WELCOME": "1"}},
		{"dumb terminal", map[string]string{"TERM": "dumb"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.env {
				t.Setenv(k, v)
			}
			if shouldShowWelcome(stdin, stdout) {
				t.Errorf("expected welcome to be disabled when %s", tc.name)
			}
		})
	}
}

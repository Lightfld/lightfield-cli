package main

import (
	"os"
	"testing"
)

// fakeFd returns a fixed file descriptor. When the fd happens to be a real
// TTY (e.g. during interactive test runs), term.IsTerminal returns true.
// For the negative-path tests we use /dev/null fds which are never TTYs.
type fakeFd uintptr

func (f fakeFd) Fd() uintptr { return uintptr(f) }

func TestShouldShowWelcome_Disabled(t *testing.T) {
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		t.Fatalf("open devnull: %v", err)
	}
	defer devNull.Close()
	fd := fakeFd(devNull.Fd())

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
			if shouldShowWelcome(fd, fd) {
				t.Errorf("expected welcome to be disabled when %s", tc.name)
			}
		})
	}
}

func TestShouldShowWelcome_NonTTY(t *testing.T) {
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		t.Fatalf("open devnull: %v", err)
	}
	defer devNull.Close()
	fd := fakeFd(devNull.Fd())

	if shouldShowWelcome(fd, fd) {
		t.Error("expected false for non-TTY file descriptors")
	}
}

func TestShouldShowWelcome_PositivePath(t *testing.T) {
	// Open /dev/tty which is the controlling terminal. Only available in
	// interactive sessions — skip in CI where there's no controlling TTY.
	tty, err := os.Open("/dev/tty")
	if err != nil {
		t.Skip("no controlling terminal available (likely CI)")
	}
	defer tty.Close()

	if !shouldShowWelcome(tty, tty) {
		t.Error("expected true for real TTY with no suppression env vars")
	}
}

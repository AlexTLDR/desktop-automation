package automation

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

// TypeString types the given text using robotgo.TypeStr
// Returns an error if the text is empty or whitespace only
func TypeString(text string) error {
	// Safety check for empty strings
	if strings.TrimSpace(text) == "" {
		return fmt.Errorf("text cannot be empty or whitespace only")
	}

	// Type the string using robotgo
	robotgo.TypeStr(text)
	return nil
}

// TypeStringWithDelay types the given text with a delay between each character
// delayMs specifies the delay in milliseconds between each character
// Returns an error if the text is empty or whitespace only, or if delayMs is negative
func TypeStringWithDelay(text string, delayMs int) error {
	// Safety check for empty strings
	if strings.TrimSpace(text) == "" {
		return fmt.Errorf("text cannot be empty or whitespace only")
	}

	// Validate delay parameter
	if delayMs < 0 {
		return fmt.Errorf("delay cannot be negative, got %d", delayMs)
	}

	// If no delay is specified, use the regular TypeString function
	if delayMs == 0 {
		return TypeString(text)
	}

	// Type each character with delay
	for _, char := range text {
		robotgo.TypeStr(string(char))
		if delayMs > 0 {
			time.Sleep(time.Duration(delayMs) * time.Millisecond)
		}
	}

	return nil
}

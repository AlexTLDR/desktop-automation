package commands

import (
	"fmt"
	"strings"

	"github.com/AlexTLDR/desktop-automation/internal/automation"

	"github.com/spf13/cobra"
)

// NewTypeCmd creates a new type command
func NewTypeCmd() *cobra.Command {
	var delayMs int

	var typeCmd = &cobra.Command{
		Use:   "type [text]",
		Short: "Type text at current cursor position",
		Long: `Type text at the current cursor position.

This command simulates keyboard input by typing the specified text at the current cursor position.
The text will be typed as if entered from the keyboard, including support for special characters.
Use quotes for multi-word text.`,
		Example: `  # Type a simple message
  desktop-automation type "Hello, World!"

  # Type text with spaces (use quotes)
  desktop-automation type "This is a test message"

  # Type a single word
  desktop-automation type hello

  # Type 1 character per second (1000ms = 1 second)
  desktop-automation type --delay=1000 "hello world"

  # Type with 500ms delay between characters
  desktop-automation type --delay=500 "Slow typing!"

  # Alternative delay flag syntax
  desktop-automation type --delay 250 "Even slower!"

  # Type text with special characters
  desktop-automation type "user@example.com"`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runType(args, delayMs)
		},
	}

	// Add delay flag with default value of 0ms
	typeCmd.Flags().IntVar(&delayMs, "delay", 0, "Delay in milliseconds between each character (default: 0)")

	return typeCmd
}

// runType executes the type command
func runType(args []string, delayMs int) error {
	// Join all arguments into a single string to support spaces
	text := strings.Join(args, " ")

	// Validate that text is not empty (this is also checked in automation functions)
	if strings.TrimSpace(text) == "" {
		return fmt.Errorf("text cannot be empty")
	}

	var err error

	// Use appropriate typing function based on delay setting
	if delayMs > 0 {
		err = automation.TypeStringWithDelay(text, delayMs)
	} else {
		err = automation.TypeString(text)
	}

	if err != nil {
		return fmt.Errorf("failed to type text: %w", err)
	}

	// Show success message with character count
	charCount := len(text)
	if delayMs > 0 {
		fmt.Printf("Successfully typed %d characters with %dms delay: %q\n", charCount, delayMs, text)
	} else {
		fmt.Printf("Successfully typed %d characters: %q\n", charCount, text)
	}

	return nil
}

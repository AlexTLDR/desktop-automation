package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// NewTypeCmd creates a new type command
func NewTypeCmd() *cobra.Command {
	var typeCmd = &cobra.Command{
		Use:   "type [text]",
		Short: "Type text at current cursor position",
		Long: `Type text at the current cursor position.

This command simulates keyboard input by typing the specified text at the current cursor position.
The text will be typed as if entered from the keyboard, including support for special characters.`,
		Example: `  # Type a simple message
  desktop-automation type "Hello, World!"

  # Type text with spaces (use quotes)
  desktop-automation type "This is a test message"

  # Type a single word
  desktop-automation type hello

  # Type text with special characters
  desktop-automation type "user@example.com"`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runType(args)
		},
	}

	return typeCmd
}

// runType executes the type command
func runType(args []string) error {
	// Join all arguments into a single string to support spaces
	text := strings.Join(args, " ")

	// Validate that text is not empty
	if strings.TrimSpace(text) == "" {
		return fmt.Errorf("text cannot be empty")
	}

	// TODO: Implement actual typing functionality
	// This is a placeholder for the actual implementation
	fmt.Printf("Typing text: %q\n", text)

	// Placeholder for actual keyboard typing implementation
	// Example: return automation.TypeText(text)

	return nil
}

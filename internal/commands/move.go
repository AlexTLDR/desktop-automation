package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// NewMoveCmd creates a new move command
func NewMoveCmd() *cobra.Command {
	var moveCmd = &cobra.Command{
		Use:   "move [x] [y]",
		Short: "Move the mouse cursor to coordinates",
		Long: `Move the mouse cursor to specific screen coordinates.

This command moves the mouse cursor to the specified x and y coordinates on the screen.
The coordinates are measured in pixels from the top-left corner of the screen (0,0).
The cursor will smoothly move to the target position without clicking.`,
		Example: `  # Move cursor to coordinates (100, 200)
  desktop-automation move 100 200

  # Move cursor to the top-left corner
  desktop-automation move 0 0

  # Move cursor to the center of a 1920x1080 screen
  desktop-automation move 960 540

  # Move cursor to coordinates (800, 600)
  desktop-automation move 800 600`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runMove(args)
		},
	}

	return moveCmd
}

// runMove executes the move command
func runMove(args []string) error {
	// Parse x coordinate
	x, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid x coordinate '%s': must be a number", args[0])
	}

	// Parse y coordinate
	y, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("invalid y coordinate '%s': must be a number", args[1])
	}

	// Validate coordinates are non-negative
	if x < 0 {
		return fmt.Errorf("x coordinate must be non-negative, got %d", x)
	}
	if y < 0 {
		return fmt.Errorf("y coordinate must be non-negative, got %d", y)
	}

	// TODO: Implement actual mouse move functionality
	// This is a placeholder for the actual implementation
	fmt.Printf("Moving mouse cursor to coordinates (%d, %d)\n", x, y)

	// Placeholder for actual mouse movement implementation
	// Example: return automation.MoveTo(x, y)

	return nil
}
